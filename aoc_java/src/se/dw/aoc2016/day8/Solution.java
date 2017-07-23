package se.dw.aoc2016.day8;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * My solution to the day8 problem
 * http://adventofcode.com/2016/day/8
 */
public class Solution {

    public Solution(String input) {

        Screen screen = new Screen(50, 6);
        for (String command : input.split("\n")) {
            screen.parseCommand(command);
        }
        screen.printScreen();

        int litPixels = screen.getLitPixels();

        System.out.println("Solution: " + litPixels);

    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/2016/day8.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}