package se.dw.aoc2016.day9;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    String input1 = "ADVENT";
    String input2 = "A(1x5)BC";
    String input3 = "(3x3)XYZ";
    String input4 = "A(2x2)BCD(2x2)EFG";
    String input5 = "(6x1)(1x3)A";
    String input6 = "X(8x2)(3x3)ABCY";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testSolution() {

        assertEquals(new Decompressor(input1).decompressData().length(), 6);
        assertEquals(new Decompressor(input2).decompressData().length(), 7);
        assertEquals(new Decompressor(input3).decompressData().length(), 9);
        assertEquals(new Decompressor(input4).decompressData().length(), 11);
        assertEquals(new Decompressor(input5).decompressData().length(), 6);
        assertEquals(new Decompressor(input6).decompressData().length(), 18);

    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}