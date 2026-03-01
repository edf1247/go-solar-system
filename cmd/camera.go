package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const (
    SPEED = 5.0
    KEY_W = 87
    KEY_S = 83
    KEY_A = 65
    KEY_D = 68
)

func InitCamera() rl.Camera3D{
    cameraPos := rl.Vector3{0, 0, 50}
    cameraTarget := rl.Vector3{0, 0, 0}
    rotation := rl.Vector3{0, 1, 0}
    fov := float32(90.0)
    camera := rl.NewCamera3D(cameraPos, cameraTarget, rotation, fov, rl.CameraPerspective)
    return camera
}
