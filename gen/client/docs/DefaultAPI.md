# \DefaultAPI

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorsActorIdDelete**](DefaultAPI.md#ActorsActorIdDelete) | **Delete** /actors/{actorId} | Удаление актёра
[**ActorsActorIdMoviesGet**](DefaultAPI.md#ActorsActorIdMoviesGet) | **Get** /actors/{actorId}/movies | Получение списка фильмов с участнием актера
[**ActorsActorIdPatch**](DefaultAPI.md#ActorsActorIdPatch) | **Patch** /actors/{actorId} | Изменение информации об актёре
[**ActorsPost**](DefaultAPI.md#ActorsPost) | **Post** /actors | Добавление актёра
[**MoviesGet**](DefaultAPI.md#MoviesGet) | **Get** /movies | Получение списка фильмов с сортировкой и поиском
[**MoviesMovieIdDelete**](DefaultAPI.md#MoviesMovieIdDelete) | **Delete** /movies/{movieId} | Удаление фильма
[**MoviesMovieIdPatch**](DefaultAPI.md#MoviesMovieIdPatch) | **Patch** /movies/{movieId} | Частичное обновление информации о фильме
[**MoviesPost**](DefaultAPI.md#MoviesPost) | **Post** /movies | Добавление фильма
[**RegisterPost**](DefaultAPI.md#RegisterPost) | **Post** /register | Регистрация пользователя и выдача токенов
[**TokenGet**](DefaultAPI.md#TokenGet) | **Get** /token | Получение информации о текущем токене
[**TokenRefreshPost**](DefaultAPI.md#TokenRefreshPost) | **Post** /token/refresh | Обновление токена доступа



## ActorsActorIdDelete

> ActorsActorIdDelete(ctx, actorId).Execute()

Удаление актёра



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	actorId := "actorId_example" // string | Уникальный идентификатор актёра

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ActorsActorIdDelete(context.Background(), actorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActorsActorIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Уникальный идентификатор актёра | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorsActorIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorsActorIdMoviesGet

> []string ActorsActorIdMoviesGet(ctx, actorId).Execute()

Получение списка фильмов с участнием актера



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	actorId := "actorId_example" // string | Уникальный идентификатор актёра

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ActorsActorIdMoviesGet(context.Background(), actorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActorsActorIdMoviesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorsActorIdMoviesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ActorsActorIdMoviesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Уникальный идентификатор актёра | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorsActorIdMoviesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorsActorIdPatch

> ActorsActorIdPatch(ctx, actorId).Actor(actor).Execute()

Изменение информации об актёре



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	actorId := "actorId_example" // string | Уникальный идентификатор актёра
	actor := *openapiclient.NewActor("mirustal", "male", time.Now()) // Actor | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ActorsActorIdPatch(context.Background(), actorId).Actor(actor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActorsActorIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Уникальный идентификатор актёра | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorsActorIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 

### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorsPost

> ActorsPost(ctx).Actor(actor).Execute()

Добавление актёра



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	actor := *openapiclient.NewActor("mirustal", "male", time.Now()) // Actor | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ActorsPost(context.Background()).Actor(actor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 

### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoviesGet

> []Movie MoviesGet(ctx).Sort(sort).Order(order).Title(title).ActorName(actorName).Execute()

Получение списка фильмов с сортировкой и поиском



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sort := "sort_example" // string | Критерий сортировки (по умолчанию - rating). (optional)
	order := "order_example" // string | Порядок сортировки (по умолчанию - desc). (optional)
	title := "title_example" // string | Фрагмент названия фильма для поиска. (optional)
	actorName := "actorName_example" // string | Фрагмент имени актёра для поиска. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MoviesGet(context.Background()).Sort(sort).Order(order).Title(title).ActorName(actorName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MoviesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoviesGet`: []Movie
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MoviesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoviesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Критерий сортировки (по умолчанию - rating). | 
 **order** | **string** | Порядок сортировки (по умолчанию - desc). | 
 **title** | **string** | Фрагмент названия фильма для поиска. | 
 **actorName** | **string** | Фрагмент имени актёра для поиска. | 

### Return type

[**[]Movie**](Movie.md)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoviesMovieIdDelete

> MoviesMovieIdDelete(ctx, movieId).Execute()

Удаление фильма



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	movieId := "movieId_example" // string | Уникальный идентификатор фильма

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.MoviesMovieIdDelete(context.Background(), movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MoviesMovieIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **string** | Уникальный идентификатор фильма | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoviesMovieIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoviesMovieIdPatch

> MoviesMovieIdPatch(ctx, movieId).MoviesMovieIdPatchRequest(moviesMovieIdPatchRequest).Execute()

Частичное обновление информации о фильме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	movieId := "movieId_example" // string | Уникальный идентификатор фильма
	moviesMovieIdPatchRequest := *openapiclient.NewMoviesMovieIdPatchRequest() // MoviesMovieIdPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.MoviesMovieIdPatch(context.Background(), movieId).MoviesMovieIdPatchRequest(moviesMovieIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MoviesMovieIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **string** | Уникальный идентификатор фильма | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoviesMovieIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moviesMovieIdPatchRequest** | [**MoviesMovieIdPatchRequest**](MoviesMovieIdPatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoviesPost

> MoviesPost(ctx).Movie(movie).Execute()

Добавление фильма



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	movie := *openapiclient.NewMovie("mirustal live", "как же он устал", time.Now(), float32(10)) // Movie | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.MoviesPost(context.Background()).Movie(movie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MoviesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoviesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movie** | [**Movie**](Movie.md) |  | 

### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPost

> RegisterPost(ctx).RegisterPostRequest(registerPostRequest).AdminMode(adminMode).Execute()

Регистрация пользователя и выдача токенов

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	registerPostRequest := *openapiclient.NewRegisterPostRequest("Username_example", "Password_example") // RegisterPostRequest | 
	adminMode := "adminMode_example" // string | Выдача админм мода (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RegisterPost(context.Background()).RegisterPostRequest(registerPostRequest).AdminMode(adminMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerPostRequest** | [**RegisterPostRequest**](RegisterPostRequest.md) |  | 
 **adminMode** | **string** | Выдача админм мода | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenGet

> TokenGet(ctx).Execute()

Получение информации о текущем токене

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.TokenGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokenGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenRefreshPost

> TokenRefreshPost(ctx).Execute()

Обновление токена доступа

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.TokenRefreshPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TokenRefreshPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokenRefreshPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[CookieAuthJWT](../README.md#CookieAuthJWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

