class BowlingGame
  def initialize
    @score = 0
    @before_shot = 0
    @is_spare = false
  end

  def roll(pins)
    @score += pins
    if @is_spare
      @score += pins
      @is_spare = false
    end
    @is_spare = true if @before_shot + pins == 10
    @before_shot = pins
  end

  def score
    @score
  end
end