package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 2560
const HEIGHT = 1600


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

        sun := CreatePlanet(rl.Vector3{0, 0, 0}, float32(10.0), float32(1000.0), rl.Yellow, "Sun", rl.Vector3{0, 0, 0})
        planet := CreatePlanet(rl.Vector3{50, 0, 0}, float32(1.0), float32(10.0), rl.Blue, "Planet", rl.Vector3{0, 20, 0})
        planet2 := CreatePlanet(rl.Vector3{-100, 0, 0}, float32(2.5), float32(50.0), rl.Red, "Planet", rl.Vector3{0, 20, 0})    

        sceneObjects = append(sceneObjects, &sun)
        sceneObjects = append(sceneObjects, &planet)
        sceneObjects = append(sceneObjects, &planet2)
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
