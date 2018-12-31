defmodule Chop do
  def binary_search(value, numbers) do
    search(0, length(numbers)-1, value, numbers)
  end

  def search(left_side, right_side, value, numbers) when left_side <= right_side do
    middle = Enum.at(numbers, find_middle(left_side, right_side))
    search(left_side, right_side, value, numbers, where_is_value(middle, value))
  end

  def search(_left_side, _right_side, _value, _numbers) do
    -1
  end

  def search(left_side, right_side, value, numbers, :left) do
    search(left_side, find_middle(left_side, right_side)-1, value, numbers)
  end

  def search(left_side, right_side, value, numbers, :right) do
    search(find_middle(left_side, right_side)+1, right_side, value, numbers)
  end

  def search(left_side, right_side, _value, _numbers, :success) do
    find_middle(left_side, right_side)
  end

  def where_is_value(middle, value) when middle > value do
    :left
  end

  def where_is_value(middle, value) when middle < value do
    :right
  end

  def where_is_value(middle, value) when middle == value do
    :success
  end

  def find_middle(left_side, right_side) do
    trunc(Float.floor((left_side + right_side) / 2))
  end
end
