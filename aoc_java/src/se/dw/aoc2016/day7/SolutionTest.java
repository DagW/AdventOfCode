package se.dw.aoc2016.day7;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    static final String validInput0 = "abba[mnop]qrst";
    static final String validInput1 = "ioxxoj[asdfgh]zxcvbn";

    static final String invalidInput0 = "abcd[bddb]xyyx";
    static final String invalidInput1 = "aaaa[qwer]tyui";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testSolution() {

        assertEquals(new IPv7Address(validInput0).supportsTLS(), true);
        assertEquals(new IPv7Address(validInput1).supportsTLS(), true);
        assertEquals(new IPv7Address(invalidInput0).supportsTLS(), false);
        assertEquals(new IPv7Address(invalidInput1).supportsTLS(), false);

    }

    @Test
    public void testAbbaLocator() {

        IPv7Address address = new IPv7Address("");
        assertEquals(address.containsAbba("abba"), true);
        assertEquals(address.containsAbba("test"), false);
        assertEquals(address.containsAbba("aaaa"), false);

    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}