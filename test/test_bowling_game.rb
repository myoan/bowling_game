require "test_helper"
require "./bowling_game"

class TestBowlingGame < Minitest::Test
  def test_all_roll_are_gutter
    game = BowlingGame.new
    20.times { game.roll(0) }
    assert_equal game.score, 0
  end

  def test_all_roll_are_1
    game = BowlingGame.new
    20.times { game.roll(1) }
    assert_equal game.score, 20
  end

  def test_simple_spare
    game = BowlingGame.new
    # 14 + 4 = 18
    game.roll(1)
    game.roll(9)
    game.roll(4)
    assert_equal game.score, 18
  end
end
