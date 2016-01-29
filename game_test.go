package bowling_test

import (
	. "github.com/nomkhonwaan/bowling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var (
		game *Game
	)

	BeforeEach(func() {
		game = NewGame()
	})

	Describe("The Bowling Game", func() {
		Context("roll 0 pin all frame", func() {
			It("should return 0 score", func() {
				for i := 0; i < 20; i++ {
					game.Roll(0)
				}

				Expect(game.Score()).To(Equal(0))
			})
		})

		Context("roll 1 pin all frame", func() {
			It("should return 20 score", func() {
				for i := 0; i < 20; i++ {
					game.Roll(1)
				}

				Expect(game.Score()).To(Equal(20))
			})
		})

		Context("spare at 1st frame and 2 score at next roll", func() {
			It("should return 14 score", func() {
				game.Roll(7)
				game.Roll(3) // spare
				game.Roll(2)
				for i := 0; i < 17; i++ {
					game.Roll(0)
				}

				Expect(game.Score()).To(Equal(14))
			})
		})

		Context("spare at 2nd frame and 3 score at next roll and 1 score in another roll", func() {
			It("should return 33 score", func() {
				game.Roll(1) // frame 1
				game.Roll(1) // frame 1
				game.Roll(5)
				game.Roll(5) // spare
				game.Roll(3)
				for i := 0; i < 15; i++ {
					game.Roll(1)
				}

				Expect(game.Score()).To(Equal(33))
			})
		})

		Context("spare at 10th frame and 4 score at next roll", func() {
			It("should return 14 score", func() {
				for i := 0; i < 18; i++ {
					game.Roll(0)
				}
				game.Roll(4)
				game.Roll(6) // spare
				game.Roll(4)

				Expect(game.Score()).To(Equal(14))
			})
		})

		Context("strike at 1st frame and 2, 4 score at next roll", func() {
			It("should return 22 score", func() {
				game.Roll(10) // strike
				game.Roll(2)
				game.Roll(4)
				for i := 0; i < 16; i++ {
					game.Roll(0)
				}

				Expect(game.Score()).To(Equal(22))
			})
		})

		Context("strike at 10th frame and make 4, 5 score at next roll", func() {
			It("should return 19 score", func() {
				for i := 0; i < 18; i++ {
					game.Roll(0)
				}
				game.Roll(10) // strike
				game.Roll(4)
				game.Roll(5)

				Expect(game.Score()).To(Equal(19))
			})
		})

		Context("strike all frame", func() {
			It("should return 300 score", func() {
				for i := 0; i < 12; i++ {
					game.Roll(10)
				}

				Expect(game.Score()).To(Equal(300))
			})
		})
	})
})
