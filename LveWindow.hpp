//
// Created by welco on 9/23/2022.
//

#ifndef CPPVULKANSTUDY_LVEWINDOW_HPP
#define CPPVULKANSTUDY_LVEWINDOW_HPP

#include <string>
#include <GLFW/glfw3.h>

namespace lve {

    class LveWindow {
    public:
        LveWindow(int w, int h, std::string name);
        ~LveWindow();

        LveWindow(const LveWindow &) = delete;
        LveWindow &operator=(const LveWindow &) = delete;

        bool shouldClose() { return glfwWindowShouldClose(window); }
    private:
        void initWindow();

        const int width;
        const int height;
        std::string windowName;
        GLFWwindow *window;
    };

} // lve

#endif //CPPVULKANSTUDY_LVEWINDOW_HPP
