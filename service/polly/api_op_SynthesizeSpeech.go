// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes. SSML input
// must be valid, well-formed SSML. Some alphabets might not be available with all
// the voices (for example, Cyrillic might not be read at all by English voices)
// unless phoneme mapping is used. For more information, see How it Works (https://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html)
// .
func (c *Client) SynthesizeSpeech(ctx context.Context, params *SynthesizeSpeechInput, optFns ...func(*Options)) (*SynthesizeSpeechOutput, error) {
	if params == nil {
		params = &SynthesizeSpeechInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SynthesizeSpeech", params, optFns, c.addOperationSynthesizeSpeechMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SynthesizeSpeechOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SynthesizeSpeechInput struct {

	// The format in which the returned output will be encoded. For audio stream, this
	// will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json. When pcm
	// is used, the content returned is audio/pcm in a signed 16-bit, 1 channel (mono),
	// little-endian format.
	//
	// This member is required.
	OutputFormat types.OutputFormat

	// Input text to synthesize. If you specify ssml as the TextType , follow the SSML
	// format for the input text.
	//
	// This member is required.
	Text *string

	// Voice ID to use for the synthesis. You can get a list of available voice IDs by
	// calling the DescribeVoices (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation.
	//
	// This member is required.
	VoiceId types.VoiceId

	// Specifies the engine ( standard or neural ) for Amazon Polly to use when
	// processing input text for speech synthesis. For information on Amazon Polly
	// voices and which voices are available in standard-only, NTTS-only, and both
	// standard and NTTS formats, see Available Voices (https://docs.aws.amazon.com/polly/latest/dg/voicelist.html)
	// . NTTS-only voices When using NTTS-only voices such as Kevin (en-US), this
	// parameter is required and must be set to neural . If the engine is not
	// specified, or is set to standard , this will result in an error. Type: String
	// Valid Values: standard | neural Required: Yes Standard voices For standard
	// voices, this is not required; the engine parameter defaults to standard . If the
	// engine is not specified, or is set to standard and an NTTS-only voice is
	// selected, this will result in an error.
	Engine types.Engine

	// Optional language code for the Synthesize Speech request. This is only
	// necessary if using a bilingual voice, such as Aditi, which can be used for
	// either Indian English (en-IN) or Hindi (hi-IN). If a bilingual voice is used and
	// no language code is specified, Amazon Polly uses the default language of the
	// bilingual voice. The default language for any voice is the one returned by the
	// DescribeVoices (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation for the LanguageCode parameter. For example, if no language code is
	// specified, Aditi will use Indian English rather than Hindi.
	LanguageCode types.LanguageCode

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon is
	// the same as the language of the voice. For information about storing lexicons,
	// see PutLexicon (https://docs.aws.amazon.com/polly/latest/dg/API_PutLexicon.html)
	// .
	LexiconNames []string

	// The audio frequency specified in Hz. The valid values for mp3 and ogg_vorbis
	// are "8000", "16000", "22050", and "24000". The default value for standard voices
	// is "22050". The default value for neural voices is "24000". Valid values for pcm
	// are "8000" and "16000" The default value is "16000".
	SampleRate *string

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []types.SpeechMarkType

	// Specifies whether the input text is plain text or SSML. The default value is
	// plain text. For more information, see Using SSML (https://docs.aws.amazon.com/polly/latest/dg/ssml.html)
	// .
	TextType types.TextType

	noSmithyDocumentSerde
}

type SynthesizeSpeechOutput struct {

	// Stream containing the synthesized speech.
	AudioStream io.ReadCloser

	// Specifies the type audio stream. This should reflect the OutputFormat parameter
	// in your request.
	//   - If you request mp3 as the OutputFormat , the ContentType returned is
	//   audio/mpeg.
	//   - If you request ogg_vorbis as the OutputFormat , the ContentType returned is
	//   audio/ogg.
	//   - If you request pcm as the OutputFormat , the ContentType returned is
	//   audio/pcm in a signed 16-bit, 1 channel (mono), little-endian format.
	//   - If you request json as the OutputFormat , the ContentType returned is
	//   application/x-json-stream.
	ContentType *string

	// Number of characters synthesized.
	RequestCharacters int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSynthesizeSpeechMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSynthesizeSpeech{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSynthesizeSpeech{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSynthesizeSpeechResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSynthesizeSpeechValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSynthesizeSpeech(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSynthesizeSpeech(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "SynthesizeSpeech",
	}
}

type opSynthesizeSpeechResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSynthesizeSpeechResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSynthesizeSpeechResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "polly"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "polly"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("polly")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addSynthesizeSpeechResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSynthesizeSpeechResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
