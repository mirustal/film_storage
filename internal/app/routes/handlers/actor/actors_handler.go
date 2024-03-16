package actors


// ActorsHandler структура для хэндлеров, связанных с актерами.
type ActorsHandler struct {
    service DefaultAPIServicer
    errorHandler ErrorHandler
}

// NewActorsHandler создает новый экземпляр ActorsHandler.
func NewActorsHandler(s DefaultAPIServicer, errorHandler ErrorHandler) *ActorsHandler {
    return &ActorsHandler{
        service: s,
        errorHandler: errorHandler,
    }
}

// Routes реализует интерфейс Handler и возвращает маршруты для актеров.
func (h *ActorsHandler) Routes() Routes {
    return Routes{
        // Определение маршрутов для актеров
    }
}