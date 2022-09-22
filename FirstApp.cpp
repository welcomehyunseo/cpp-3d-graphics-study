//
// Created by welco on 9/23/2022.
//
#include "FirstApp.hpp"

namespace lve {
    void FirstApp::run() {
        while (!lveWindow.shouldClose()) {
            glfwPollEvents();
        }
    }
}