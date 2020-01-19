Feature: positioning a robot
  In order to position a robot in the plateau
  As an instructor
  I need to be send correct instructions

  Scenario Outline: Send request to position a robot in the mars plateau
    Given the layout of the plateau is determined
    When I send a request "<details>" to position the robot within the plateau boundaries
    And  I send a set of "<commands>" instructions to the robot
    Then I am able to see the final "<output>" position of the robot in the plateau

    Examples:
      | details |  | commands |  | output |
      | 1 2 N |    | LMLMLMLMM | |  1 3 N |


