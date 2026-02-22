package main

import rl "github.com/gen2brain/raylib-go/raylib"

const WIDTH = 800
const HEIGHT = 450

func main() {
  rl.InitWindow(WIDTH, HEIGHT, "Go Engine")
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)
  var camInit = false
  var mainCamera rl.Camera3D

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(BLACK)
    
    if !camInit {
        mainCamera = InitCamera()
        camInit = true
    }
    rl.UpdateCamera(&mainCamera, rl.CameraFree)
    rl.BeginMode3D(mainCamera)
    
    moveCamera(&mainCamera)

    rl.DrawCube(rl.Vector3{0, 0, 0}, 2.0, 2.0, 2.0, WHITE)
    rl.DrawCubeWires(rl.Vector3{0, 0, 0}, 2.0, 2.0, 2.0, MAROON)
    rl.DrawGrid(10.0, 1.0)

    rl.EndMode3D()
    rl.EndDrawing()
  }
}
