require "test_helper"

class TestBowlingGame < Minitest::Test
  def test_simple_roll
    game = BowlingGame.new
    game.roll(0)
    assert_equal game.score, 0
  end
end
