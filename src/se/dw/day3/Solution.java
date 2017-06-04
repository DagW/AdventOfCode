package se.dw.day3;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * My solution to the day3 problem
 * http://adventofcode.com/2016/day/3
 */
public class Solution {

    public Solution(String input) {

        /*
        How many of the input lines are valid triangles
        One line contains the lenght of the three sides
        The sum of any two numbers must be larger than the remaining one
         */

        int validTriangleCount = 0;

        for (String triangleInput : input.split("\n")) {

            Triangle triangle = new Triangle(triangleInput);

            if (triangle.isValid()) {
                validTriangleCount++;
            }

        }

        System.out.println("Solution: " + validTriangleCount);

    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/day3.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}