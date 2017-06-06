package se.dw.day6;

import org.junit.Test;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import static org.junit.Assert.assertEquals;

public class SolutionTest {


    public final static String shouldBeEaster = "eedadn\n" +
            "drvtee\n" +
            "eandsr\n" +
            "raavrd\n" +
            "atevrs\n" +
            "tsrnev\n" +
            "sdttsa\n" +
            "rasrtv\n" +
            "nssdts\n" +
            "ntnada\n" +
            "svetve\n" +
            "tesnvt\n" +
            "vntsnd\n" +
            "vrdear\n" +
            "dvrsen\n" +
            "enarar";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testSolution() throws IOException {

        Path path = Paths.get("res/day6.txt");
        String input = new String(Files.readAllBytes(path));
        assertEquals(new Solution(input).getSolution(), "tsreykjj");

        assertEquals(new Solution(shouldBeEaster).getSolution(), "easter");

    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}