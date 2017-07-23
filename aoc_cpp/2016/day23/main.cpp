#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <map>

/*
 * https://stackoverflow.com/questions/2844817/how-do-i-check-if-a-c-string-is-an-int
 */
inline bool isInteger(const std::string &s) {
    if (s.empty() || ((!isdigit(s[0])) && (s[0] != '-') && (s[0] != '+'))) return false;

    char *p;
    std::strtol(s.c_str(), &p, 10);

    return (*p == 0);
}

class Registers {

private:
    std::map<std::string, int> registers;

public:

    Registers() {
        registers["a"] = 7; //Day 23 special: init a => 7
        registers["b"] = 0;
        registers["c"] = 0;
        registers["d"] = 0;
    }

    void cpy(std::string input, std::string toRegister) {
        if (isInteger(input)) {
            registers[toRegister] = std::stoi(input);
        } else {
            registers[toRegister] = registers[input];
        }
    }

    void inc(std::string reg) {
        registers[reg] = registers[reg] + 1;
    }

    void dec(std::string reg) {
        registers[reg] = registers[reg] - 1;
    }

    void show() {
        std::cout << "Registers" << std::endl;
        for (auto elem : registers) {
            std::cout << elem.first << " " << elem.second << "\n";
        }
    }

    int getValue(std::string reg) {
        return registers[reg];
    }
};


int main() {

    std::vector<std::vector<std::string>> rows;

    std::ifstream file("2016/day23/input.txt");
    for (std::string line; getline(file, line);) {

        std::vector<std::string> instructionParts;

        std::istringstream f(line);
        std::string column;

        while (std::getline(f, column, ' ')) {
            instructionParts.push_back(column);
        }

        rows.push_back(instructionParts);
    }

    Registers registers;
    for (int i = 0; i < rows.size(); ++i) {

        std::string instruction = rows[i][0];

        if (instruction == "cpy") {
            registers.cpy(rows[i][1], rows[i][2]);
        } else if (instruction == "inc") {
            registers.inc(rows[i][1]);
        } else if (instruction == "dec") {
            registers.dec(rows[i][1]);
        } else if (instruction == "jnz") {

            std::string reg = rows[i][1];
            std::string relativeJump = rows[i][2];

            //The reg cant be "0" if its numeric
            //And if its a register the content cant be 0
            if ((isInteger(reg) && reg != "0") ||
                registers.getValue(reg) != 0) {

                //Move the position forward, relative to current iteration
                if (isInteger(relativeJump)) {
                    i += std::stoi(relativeJump) - 1;
                } else {
                    i += registers.getValue(relativeJump) - 1;
                }
            }

        } else if (instruction == "tgl") {
            int relativeToggle = registers.getValue(rows[i][1]);
            int index = i + relativeToggle;

            if (index >= rows.size()) {
                //Nothing
            } else if (index == i) {
                //If its the current position, become inc
                rows[index][0] = "inc";
            } else if (rows[index][0] == "inc") {
                //Inc becomes dec
                rows[index][0] = "dec";
            } else if (rows[index][0] == "jnz") {
                //Jnz becomes cpy
                rows[index][0] = "cpy";
            } else if (rows[index].size() == 2) {
                //One arg instructions
                rows[index][0] = "inc";
            } else if (rows[index].size() == 3) {
                //Two arg instructions
                rows[index][0] = "jnz";
            }

        }
    }
    registers.show();

    return 0;

}
