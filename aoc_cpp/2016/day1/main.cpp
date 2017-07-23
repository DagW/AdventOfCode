#include <iostream>
#include <fstream>

int direction = 0; //Start facing north
int current_position[2] = {0, 0}; //At position x,y=0,0

void walk(int distance) {

    switch (direction) {
        case 0:
            current_position[1] += distance; //North
            break;
        case 1:
            current_position[0] += distance; //East
            break;
        case 2:
            current_position[1] -= distance; //South
            break;
        case 3:
            current_position[0] -= distance; // West
            break;
    }
    //std::cout << "Walking " << distance << " in direction " << direction << std::endl;

}

void turn(std::string inDirection) {

    if (inDirection == "L") {
        direction--;
    } else if (inDirection == "R") {
        direction++;
    }

    if (direction < 0) {
        direction = 3;
    }

    direction %= 4;
    //std::cout << inDirection << " -> "<< direction << std::endl;

}

int main() {

    std::ifstream file("2016/day1/input.txt");
    std::string instruction;

    while (getline(file, instruction, ',')) {

        instruction.erase(std::remove(instruction.begin(), instruction.end(), ' '), instruction.end());

        std::string direction = instruction.substr(0, 1);
        turn(direction);

        int distance = std::stoi(instruction.substr(1, instruction.length()));
        walk(distance);

    }

    std::cout << "Current position: (x,y)=(" << current_position[0] << "," << current_position[1] << ")" << std::endl;

    int solution = std::abs(current_position[0]) + std::abs(current_position[1]);
    std::cout << "Solution: " << solution << std::endl;

    return 0;

}
