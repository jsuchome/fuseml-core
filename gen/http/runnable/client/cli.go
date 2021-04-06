// Code generated by goa v3.3.1, DO NOT EDIT.
//
// runnable HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fuseml/fuseml-core/design

package client

import (
	"encoding/json"
	"fmt"

	runnable "github.com/fuseml/fuseml-core/gen/runnable"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the runnable list endpoint from CLI
// flags.
func BuildListPayload(runnableListID string) (*runnable.ListPayload, error) {
	var id *string
	{
		if runnableListID != "" {
			id = &runnableListID
		}
	}
	v := &runnable.ListPayload{}
	v.ID = id

	return v, nil
}

// BuildRegisterPayload builds the payload for the runnable register endpoint
// from CLI flags.
func BuildRegisterPayload(runnableRegisterBody string) (*runnable.Runnable, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(runnableRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"created\": \"date-time\",\n      \"id\": \"uuid\",\n      \"image\": {\n         \"registryUrl\": \"Neque dolorem ullam deserunt excepturi ut.\",\n         \"repository\": \"Corporis necessitatibus cupiditate voluptatem amet neque dolorem.\",\n         \"tag\": \"Qui corrupti aliquid voluptatem architecto inventore.\"\n      },\n      \"inputs\": [\n         {\n            \"kind\": \"Distinctio magni quisquam aut temporibus vel accusantium.\",\n            \"name\": \"Autem ipsa voluptas ut commodi quidem.\",\n            \"runnable\": \"Facere dolore eius.\"\n         },\n         {\n            \"kind\": \"Distinctio magni quisquam aut temporibus vel accusantium.\",\n            \"name\": \"Autem ipsa voluptas ut commodi quidem.\",\n            \"runnable\": \"Facere dolore eius.\"\n         },\n         {\n            \"kind\": \"Distinctio magni quisquam aut temporibus vel accusantium.\",\n            \"name\": \"Autem ipsa voluptas ut commodi quidem.\",\n            \"runnable\": \"Facere dolore eius.\"\n         }\n      ],\n      \"kind\": \"Fugit omnis quo distinctio blanditiis explicabo.\",\n      \"labels\": [\n         \"Fuga quia.\",\n         \"Eos facere ipsam ducimus repellat vel nam.\",\n         \"Cupiditate nulla dolor qui illum ut non.\",\n         \"Non iusto perferendis optio tempore voluptate.\"\n      ],\n      \"name\": \"Sit commodi labore vel voluptatum inventore.\",\n      \"outputs\": [\n         {\n            \"kind\": \"Totam animi.\",\n            \"metadata\": {\n               \"datatype\": \"Ea et ea voluptatibus quod.\",\n               \"default\": \"Laboriosam id et rem voluptatem laudantium aliquam.\",\n               \"optional\": false\n            },\n            \"name\": \"Et dolores molestias accusantium.\",\n            \"runnable\": {\n               \"kind\": \"Qui nostrum.\",\n               \"labels\": [\n                  \"Earum voluptas dolor.\",\n                  \"Nisi quis quas alias aut possimus.\",\n                  \"Libero omnis eum ea amet suscipit.\",\n                  \"Error itaque sequi.\"\n               ],\n               \"name\": \"Omnis natus.\"\n            }\n         },\n         {\n            \"kind\": \"Totam animi.\",\n            \"metadata\": {\n               \"datatype\": \"Ea et ea voluptatibus quod.\",\n               \"default\": \"Laboriosam id et rem voluptatem laudantium aliquam.\",\n               \"optional\": false\n            },\n            \"name\": \"Et dolores molestias accusantium.\",\n            \"runnable\": {\n               \"kind\": \"Qui nostrum.\",\n               \"labels\": [\n                  \"Earum voluptas dolor.\",\n                  \"Nisi quis quas alias aut possimus.\",\n                  \"Libero omnis eum ea amet suscipit.\",\n                  \"Error itaque sequi.\"\n               ],\n               \"name\": \"Omnis natus.\"\n            }\n         }\n      ]\n   }'")
		}
		if body.Image == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
		}
		if body.Inputs == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("inputs", "body"))
		}
		if body.Outputs == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("outputs", "body"))
		}
		if body.Labels == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("labels", "body"))
		}
		if body.ID != nil {
			err = goa.MergeErrors(err, goa.ValidatePattern("body.id", *body.ID, "uuid"))
		}
		if body.Created != nil {
			err = goa.MergeErrors(err, goa.ValidatePattern("body.created", *body.Created, "date-time"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &runnable.Runnable{
		ID:      body.ID,
		Name:    body.Name,
		Kind:    body.Kind,
		Created: body.Created,
	}
	if body.Image != nil {
		v.Image = marshalRunnableImageRequestBodyToRunnableRunnableImage(body.Image)
	}
	if body.Inputs != nil {
		v.Inputs = make([]*runnable.RunnableInput, len(body.Inputs))
		for i, val := range body.Inputs {
			v.Inputs[i] = marshalRunnableInputRequestBodyToRunnableRunnableInput(val)
		}
	}
	if body.Outputs != nil {
		v.Outputs = make([]*runnable.RunnableOutput, len(body.Outputs))
		for i, val := range body.Outputs {
			v.Outputs[i] = marshalRunnableOutputRequestBodyToRunnableRunnableOutput(val)
		}
	}
	if body.Labels != nil {
		v.Labels = make([]string, len(body.Labels))
		for i, val := range body.Labels {
			v.Labels[i] = val
		}
	}

	return v, nil
}

// BuildGetPayload builds the payload for the runnable get endpoint from CLI
// flags.
func BuildGetPayload(runnableGetRunnableNameOrID string) (*runnable.GetPayload, error) {
	var runnableNameOrID string
	{
		runnableNameOrID = runnableGetRunnableNameOrID
	}
	v := &runnable.GetPayload{}
	v.RunnableNameOrID = runnableNameOrID

	return v, nil
}