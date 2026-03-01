package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 1920
const HEIGHT = 1080


func main() {
  rl.InitWindow(WIDTH, HEIGHT, "Go Engine")
  defer rl.CloseWindow()

  var start = false
  var mainCamera rl.Camera3D
  var sm SceneManager
  var sceneObjects []*Body
  var counter = 0

  rl.SetTargetFPS(60.0)


  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)
    rl.DrawFPS(10.0, 10.0)

    if !start {
        mainCamera = InitCamera()

        radius := float32(10.0)
        mass := float32(1.0)

        sun := CreatePlanet(rl.Vector3{0, 0, 0}, radius, float32(100.0), rl.RayWhite, "Sun")
        planet := CreatePlanet(rl.Vector3{40, 0, 0}, float32(1.0), mass, rl.Blue, "Planet")

        sceneObjects = append(sceneObjects, &sun)
        sceneObjects = append(sceneObjects, &planet)
        sm.objects = &sceneObjects

        start = true
    }

    rl.UpdateCamera(&mainCamera, rl.CameraFree)
    
    rl.BeginMode3D(mainCamera)
    
    sm.Render()

    UniversalGravitation(&sm)
    UpdatePositions(&sm)

    rl.DrawGrid(10.0, 1.0)
    rl.EndMode3D()

    rl.EndDrawing()
    counter += 1
  }
}
