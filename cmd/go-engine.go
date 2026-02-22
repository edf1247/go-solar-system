package main

import (
    camera "github.com/edf1247/go-engine/internal/camera"
    colors "github.com/edf1247/go-engine/internal/colors"
    rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 800
const HEIGHT = 450

func main() {
  rl.InitWindow(WIDTH, HEIGHT, "Go Engine")
  defer rl.CloseWindow()

  rl.SetTargetFPS(144.0)
  var camInit = false
  var mainCamera rl.Camera3D

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(colors.BLACK)
    
    if !camInit {
        mainCamera = camera.InitCamera()
        camInit = true
    }
    rl.UpdateCamera(&mainCamera, rl.CameraFree)
    rl.BeginMode3D(mainCamera)
    
    camera.MoveCamera(&mainCamera)

    rl.DrawCube(rl.Vector3{0, 0, 0}, 2.0, 2.0, 2.0, colors.WHITE)
    rl.DrawCubeWires(rl.Vector3{0, 0, 0}, 2.0, 2.0, 2.0, colors.MAROON)
    rl.DrawGrid(10.0, 1.0)

    rl.EndMode3D()
    rl.EndDrawing()
  }
}
