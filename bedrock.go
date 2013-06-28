package bedrock

import (
	"errors"
	gl "github.com/chsc/gogl/gl33"
	"github.com/jteeuwen/glfw"
)

var ScreenWidth int
var ScreenHeight int

func Init() error {
	if err := glfw.Init(); err != nil {
		return errors.New("bedrock-init: Could not initialize glfw, said: " + err.Error())
	}

	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 0)
	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	ScreenWidth = 1280
	ScreenHeight = 720

	if err := glfw.OpenWindow(ScreenWidth, ScreenHeight, 8, 8, 8, 8, 16, 0, glfw.Windowed); err != nil {
		return errors.New("bedrock-init: Failed to open window, said: " + err.Error())
	}

	glfw.SetSwapInterval(0)
	glfw.SetWindowTitle("Dwelling")

	if err := gl.Init(); err != nil {
		return errors.New("bedrock-init: Could not initialize gl, said: " + err.Error())
	}

	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.5, 0.5, 0.5, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, gl.Sizei(ScreenWidth), gl.Sizei(ScreenHeight))

	return nil
}

func Cleanup() {
	glfw.Terminate()
	glfw.CloseWindow()
}
