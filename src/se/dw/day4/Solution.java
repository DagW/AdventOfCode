package se.dw.day4;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * My solution to the day4 problem
 * http://adventofcode.com/2016/day/4
 */
public class Solution {

    public Solution(String input) {

        int sumRoomIds = 0;

        for (String roomString : input.split("\n")) {

            Room room = new Room(roomString);
            if (room.isChecksumValid()) {
                sumRoomIds += room.getRoomId();
            }

        }

        System.out.println("Solution: " + sumRoomIds);

    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/day4.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}