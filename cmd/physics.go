package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "math"
)

type RigidBody struct {
    mass float32
    initialVelocity rl.Vector3
    currentVelocity rl.Vector3
    centerPos rl.Vector3
}

type Mesh struct {
    radius float32
    rings int32
    slices int32
    color rl.Color
}

type Body struct {
    mesh Mesh
    rb RigidBody
    tag string
}

const (
    G = math.Pi * math.Pi * 4
    VELOCITY = 10
    rings = 15
    slices = 15
)

func (planet* Body) Render() {
    rl.DrawSphereEx(planet.rb.centerPos, planet.mesh.radius, planet.mesh.rings, planet.mesh.slices, planet.mesh.color)
    
    factor := 0.75
    pColor := planet.mesh.color
    wireColor := rl.Color{uint8(float64(pColor.R) * factor * 1.1), uint8(float64(pColor.G) * factor), uint8(float64(pColor.B) * factor * 1.2), 255}
    rl.DrawSphereWires(planet.rb.centerPos, planet.mesh.radius, planet.mesh.rings, planet.mesh.slices, wireColor)
}

func CreatePlanet(centerPos rl.Vector3, radius float32, mass float32, color rl.Color, tag string, initialVelocity rl.Vector3) (Body) {
    temp := Body{Mesh{radius, rings, slices, color}, RigidBody{mass, initialVelocity, initialVelocity, centerPos}, tag}
    return temp
}

func UniversalGravitation(sm *SceneManager) {
    delta := rl.GetFrameTime()
    for _, curBody := range *sm.objects {
        for _, otherBody := range *sm.objects {
            if curBody != otherBody {
                distance := rl.Vector3DistanceSqr(curBody.rb.centerPos, otherBody.rb.centerPos)
                forceDirection := rl.Vector3Normalize(rl.Vector3Subtract(otherBody.rb.centerPos, curBody.rb.centerPos))
                force := rl.Vector3Scale(forceDirection, G * curBody.rb.mass * otherBody.rb.mass / distance)
                acceleration := rl.Vector3Scale(force, 1 / curBody.rb.mass)
                curBody.rb.currentVelocity = rl.Vector3Add(curBody.rb.currentVelocity, rl.Vector3Scale(acceleration, delta)) 
            }
        }
    }
}

func UpdatePositions(sm *SceneManager) {
    delta := rl.GetFrameTime()
    for _, curBody := range *sm.objects {
        collision := false
        for _, otherBody := range *sm.objects {
            if otherBody != curBody && rl.CheckCollisionSpheres(curBody.rb.centerPos, curBody.mesh.radius, otherBody.rb.centerPos, otherBody.mesh.radius) {
                collision = true
            }
        }
        if !collision {
            curBody.rb.centerPos = rl.Vector3Add(curBody.rb.centerPos, rl.Vector3Scale(curBody.rb.currentVelocity, delta))
        }
    }
}
