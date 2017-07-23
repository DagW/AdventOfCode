package se.dw.aoc2016.day5;

import org.junit.Test;

import java.security.NoSuchAlgorithmException;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    public final static String exampleinput = "abc";
    public final static String input = "ffykfhsq";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testSolution() throws NoSuchAlgorithmException {
        assertEquals(new Solution(exampleinput).password, "18F47A30");
        assertEquals(new Solution(input).password, "C6697B55");
    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}