package se.dw.aoc2016.day5;

import static org.junit.Assert.assertEquals;

import java.security.NoSuchAlgorithmException;
import org.junit.Test;

public class SolutionTest {

  public static final String exampleinput = "abc";
  public static final String input = "ffykfhsq";

  @org.junit.Before
  public void setUp() throws Exception {}

  @Test
  public void testSolution() throws NoSuchAlgorithmException {
    assertEquals(new Solution(exampleinput).password, "18F47A30");
    assertEquals(new Solution(input).password, "C6697B55");
  }

  @org.junit.After
  public void tearDown() throws Exception {}
}
