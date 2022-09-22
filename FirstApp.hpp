//
// Created by welco on 9/23/2022.
//

#ifndef CPPVULKANSTUDY_FIRSTAPP_HPP
#define CPPVULKANSTUDY_FIRSTAPP_HPP

#include "LveWindow.hpp"

namespace lve {
    class FirstApp {
    public:
        static constexpr int WIDTH = 800;
        static constexpr int HEIGHT = 600;

        void run();
    private:
        LveWindow lveWindow{WIDTH, HEIGHT, "Hello Vulkan!"};
    };
}

#endif //CPPVULKANSTUDY_FIRSTAPP_HPP
