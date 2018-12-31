ExUnit.start

defmodule ChopTest do
  use ExUnit.Case

  test "simple cases for binary search" do
    assert Chop.binary_search(3, []) == -1
    assert Chop.binary_search(3, [1]) == -1
    assert Chop.binary_search(1, [1]) == 0
  end

  test "with odd lenght of numbers" do
    assert Chop.binary_search(1, [1, 3, 5]) == 0
    assert Chop.binary_search(3, [1, 3, 5]) == 1
    assert Chop.binary_search(5, [1, 3, 5]) == 2
    assert Chop.binary_search(0, [1, 3, 5]) == -1
    assert Chop.binary_search(2, [1, 3, 5]) == -1
    assert Chop.binary_search(4, [1, 3, 5]) == -1
    assert Chop.binary_search(6, [1, 3, 5]) == -1
  end

  test "with even legnth of numbers" do
    assert Chop.binary_search(1, [1, 3, 5, 7]) == 0
    assert Chop.binary_search(3, [1, 3, 5, 7]) == 1
    assert Chop.binary_search(5, [1, 3, 5, 7]) == 2
    assert Chop.binary_search(7, [1, 3, 5, 7]) == 3
    assert Chop.binary_search(0, [1, 3, 5, 7]) == -1
    assert Chop.binary_search(2, [1, 3, 5, 7]) == -1
    assert Chop.binary_search(4, [1, 3, 5, 7]) == -1
    assert Chop.binary_search(6, [1, 3, 5, 7]) == -1
    assert Chop.binary_search(8, [1, 3, 5, 7]) == -1
  end
end
