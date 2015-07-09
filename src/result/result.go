package result


type Result struct {
  Success interface{}
  Failure error
}

func NewFailure(err error) Result {
  result := Result {
    Success: nil,
    Failure: err,
  }

  return result
}

func NewSuccess(value interface{}) Result {
  result := Result {
    Success: value,
    Failure: nil,
  }

  return result
}

func (result Result) Analysis(ifSuccess func(interface{}) Result, ifFailure func(error) Result) Result {
  if result.Success != nil {
    return ifSuccess(result.Success)
  }

  return ifFailure(result.Failure)
}

func (result Result) FlatMap(transform func(interface{}) Result) Result {
  if result.Success != nil {
    return transform(result.Success)
  }

  return result
}

func Try(closure func()(value interface{}, err error)) Result {
  value, err := closure()

  if err != nil {
    return NewFailure(err)
  }

  return NewSuccess(value)
}

