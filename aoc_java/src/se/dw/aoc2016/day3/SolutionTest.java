package se.dw.aoc2016.day3;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    static String validTriangle = "100 100 100";
    static String invalidTriange = "1 90 100";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testTriangle() {
        assertEquals(new Triangle(validTriangle).isValid(), true);
        assertEquals(new Triangle(invalidTriange).isValid(), false);
    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}