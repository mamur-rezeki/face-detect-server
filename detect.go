package main

import (
	"bytes"
	"errors"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/Kagami/go-face"
)

var FaceDetector *face.Recognizer

func init() {
	if fd, err := face.NewRecognizer("./models"); err != nil {
		LogCode(err)
	} else {
		FaceDetector = fd
	}

	http.HandleFunc("/", faceDetect)

}

func faceDetect(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		LogCode()

		if data, _, err := r.FormFile("foto"); err != nil {
			LogCode(err)
		} else {
			d := streamToByte(data)
			if _, dd, err := faceDetecting(d); err != nil {
				LogCode(err)
			} else {
				if dd != nil {

					w.Header().Add("Content-type", http.DetectContentType(dd))
					w.Write(dd)
				}
			}
		}

		w.Write([]byte(r.FormValue("name")))

	} else {
		var tmpl, _ = template.ParseFiles("form.html")
		tmpl.Execute(w, nil)
	}
}

func faceDetecting(data []byte) ([]ObjectDetected, []byte, error) {
	if faces, err := FaceDetector.Recognize(data); err != nil {
		return nil, nil, err
	} else {

		if len(faces) > 0 {
			if image_res, _, err := image.Decode(bytes.NewReader(data)); err != nil {
				return nil, nil, err
			} else {

				var detected []ObjectDetected
				image_dst := image.NewRGBA(image_res.Bounds())
				draw.Draw(image_dst, image_res.Bounds(), image_res, image.Point{}, draw.Src)

				for i, f := range faces {
					LogCode()

					draw.Draw(image_dst, f.Rectangle, &image.Uniform{C: color.Black}, image.Point{}, draw.Src)
					detected = append(detected, ObjectDetected{
						Id:   i,
						Rect: f.Rectangle,
					})

				}
				var rsb bytes.Buffer
				if err := jpeg.Encode(&rsb, image_dst, &jpeg.Options{Quality: 100}); err != nil {
					LogCode(err)
				}
				return detected, rsb.Bytes(), err

			}

		}
	}

	return nil, data, errors.New("can't detect faces")
}
