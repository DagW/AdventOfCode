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


    int validTriangleCount;

    public Solution(String input) {

        /*
        How many of the input lines are valid triangles
        One line contains the lenght of the three sides
        The sum of any two numbers must be larger than the remaining one
         */

        validTriangleCount = 0;

        for (String triangleInput : input.split("\n")) {

            Triangle triangle = new Triangle(triangleInput);

            if (triangle.isValid()) {
                validTriangleCount++;
            }

        }

        System.out.println("Solution: " + validTriangleCount);

    }

    private class Triangle {

        private final int side0, side1, side2;

        public Triangle(String triangle) {
            String[] stringSides = triangle.trim().split("\\s+");

            side0 = Integer.parseInt(stringSides[0]);
            side1 = Integer.parseInt(stringSides[1]);
            side2 = Integer.parseInt(stringSides[2]);
        }

        public boolean isValid() {

            //All two side sums must be larger than the third side
            if (side0 + side1 <= side2) {
                return false;
            } else if (side1 + side2 <= side0) {
                return false;
            } else if (side2 + side0 <= side1) {
                return false;
            }
            return true;

        }
    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/day3.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}