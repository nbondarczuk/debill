
type LocationOK struct {
  Body []model.Location
}

func (m *LocationOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer){
  // generated code here
}

type LocationError struct {
  Body models.Error
}

func (m *LocationError) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer){
  // generated code here
}

type LocationHandler struct {
  db interface {
    FetchLocations(*LocationParams) ([]model.Location, error)
  }
}

func (m *LocationHandler) Handle(params LocationParams) middleware.Responder {
  locations, err := m.db.FetchLocationss(&params)
  if err != nil {
    return &LocationError{Body: model.Error{Message: err.Error()}}
  }
  return &LocationOK{Body: loctionss}
}
