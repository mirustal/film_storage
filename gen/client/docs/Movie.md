# Movie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Уникальный идентификатор фильма | [optional] 
**Title** | **string** |  | 
**Description** | **string** |  | 
**ReleaseDate** | **string** |  | 
**Rating** | **float32** |  | 

## Methods

### NewMovie

`func NewMovie(title string, description string, releaseDate string, rating float32, ) *Movie`

NewMovie instantiates a new Movie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieWithDefaults

`func NewMovieWithDefaults() *Movie`

NewMovieWithDefaults instantiates a new Movie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Movie) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Movie) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Movie) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Movie) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Movie) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Movie) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Movie) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Movie) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Movie) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Movie) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReleaseDate

`func (o *Movie) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Movie) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Movie) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetRating

`func (o *Movie) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Movie) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Movie) SetRating(v float32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


