#include <iostream>
#include <fstream>


class DialPad {

private:
    int dialpad[3][3] = {
            {1, 4, 7},
            {2, 5, 8},
            {3, 6, 9}
    };
    int currentPosition[2] = {1, 1};

public:
    void moveDown() {
        currentPosition[1]++;
        if (currentPosition[1] > 2)
            currentPosition[1] = 2;
    }

    void moveUp() {
        currentPosition[1]--;
        if (currentPosition[1] < 0)
            currentPosition[1] = 0;
    }

    void moveRight() {
        currentPosition[0]++;
        if (currentPosition[0] > 2)
            currentPosition[0] = 2;
    }

    void moveLeft() {
        currentPosition[0]--;
        if (currentPosition[0] < 0)
            currentPosition[0] = 0;
    }

    void show() {
        std::cout << dialpad[currentPosition[0]][currentPosition[1]];
    }

};

int main() {

    DialPad pad;

    std::ifstream file("2016/day2/input.txt");
    std::string row;

    while (file >> row) {

        for (char &ch : row) {

            switch (ch) {
                case 'U':
                    pad.moveUp();
                    break;
                case 'D':
                    pad.moveDown();
                    break;
                case 'L':
                    pad.moveLeft();
                    break;
                case 'R':
                    pad.moveRight();
                    break;
            }

        }

        pad.show();
    }

    return 0;
}
