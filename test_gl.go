package main

import (
    //"math/rand"
    "runtime"
    //"time"
    "github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

func mainGL() {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()

    // 只能在主线程渲染，下面会报错
    //go func() {
        //tk := time.NewTicker(time.Millisecond)
        //for range tk.C {
            //width := rand.Float32()*1000 + 1
            //height := rand.Float32()*1000 + 1
            //window.SetSize(int(width), int(height))
        //}
    //}()

    for !window.ShouldClose() {
        // Do OpenGL stuff.
        window.SwapBuffers()
        glfw.PollEvents()
    }
}

