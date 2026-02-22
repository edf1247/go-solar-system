package main

import rl "github.com/gen2brain/raylib-go/raylib"

const SPEED = 5.0
const KEY_W = 87
const KEY_S = 83
const KEY_A = 65
const KEY_D = 68

func InitCamera() rl.Camera3D{
    cameraPos := rl.Vector3{10, 10, 0}
    cameraTarget := rl.Vector3{0, 0, 0}
    rotation := rl.Vector3{0, 1, 0}
    fov := float32(90.0)
    camera := rl.NewCamera3D(cameraPos, cameraTarget, rotation, fov, rl.CameraPerspective)
    return camera
}

func moveCamera(mainCam *rl.Camera3D) {
    if (rl.IsKeyDown(KEY_W)) {
        rl.CameraMoveForward(mainCam, (SPEED * rl.GetFrameTime()), 0)
    } //check if w is pressed
    if (rl.IsKeyDown(KEY_S)) {
        rl.CameraMoveForward(mainCam, -(SPEED * rl.GetFrameTime()), 0)
    }
    if (rl.IsKeyDown(KEY_A)) {
        rl.CameraMoveRight(mainCam, -(SPEED * rl.GetFrameTime()), 0)
    }
    if (rl.IsKeyDown(KEY_D)) {
        rl.CameraMoveRight(mainCam, (SPEED * rl.GetFrameTime()), 0)
    }
}
