.PHONY: all

build:
	set CGO_CXXFLAGS=--std=c++11
	set CGO_CPPFLAGS=-IC:\opencv\build\install\include
	set CGO_LDFLAGS=-LC:\opencv\build\install\x64\mingw\lib -lopencv_core454 -lopencv_face454 -lopencv_videoio454 -lopencv_imgproc454 -lopencv_highgui454 -lopencv_imgcodecs454 -lopencv_objdetect454 -lopencv_features2d454 -lopencv_video454 -lopencv_dnn454 -lopencv_xfeatures2d454 -lopencv_plot454 -lopencv_tracking454 -lopencv_img_hash454 -lopencv_calib3d454
	go build -tags static --ldflags '-extldflags="-static"' -o build/koolo.exe ./cmd/koolo/main.go
	xcopy /E /I assets build\assets
	xcopy /E /I config build\config