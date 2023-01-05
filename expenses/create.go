package expenses

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func CreateExpenses(c echo.Context) error {
	var exp Expenses
	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	row := db.QueryRow("INSERT INTO expenses (title, amount,note,tags) values ($1, $2,$3,$4) RETURNING id", exp.Title, exp.Amount, exp.Note, pq.Array(&exp.Tags))
	err = row.Scan(&exp.ID)
	if err != nil {
		fmt.Println("can't scan id", err)
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}
	fmt.Println("insert todo success id : ", exp)
	return c.JSON(http.StatusCreated, exp)

}