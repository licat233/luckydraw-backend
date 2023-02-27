package logic

import (
	"context"
	"fmt"
	"image/color"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMathCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMathCaptchaLogic {
	return &GetMathCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type CaptchaConfig struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func (l *GetMathCaptchaLogic) GetMathCaptcha(req *types.CaptchaReq) (any, error) {
	captchaConfig := &CaptchaConfig{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			Height:          req.Height,
			Width:           req.Width,
			NoiseCount:      10,
			ShowLineOptions: base64Captcha.OptionShowSlimeLine,
			BgColor:         &color.RGBA{R: req.BgColor.R, G: req.BgColor.G, B: req.BgColor.B, A: req.BgColor.A},
			Fonts:           []string{"actionj.ttf", "wqy-microhei.ttc", "chromohv.ttf"},
		},
	}

	driver := captchaConfig.DriverMath.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, l.svcCtx.CaptchaStore)
	captchaId, base64Image, err := c.Generate()
	if err != nil {
		l.Logger.Error("验证码生成失败", err)
		return nil, errorx.InternalError(err)
	}
	data := &types.CaptchaInfo{
		CaptchaId:   captchaId,
		Base64Image: base64Image,
	}
	return respx.New(data), nil
}

func GenerateDigitCaptcha(store base64Captcha.Store) (captchaId string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("panic err:%+v", e)
		}
	}()

	driver := base64Captcha.NewDriverDigit(40, 150, 6, 0, 0)
	c := base64Captcha.NewCaptcha(driver, store)
	captchaId, _, err = c.Generate()
	if err != nil {
		err = fmt.Errorf("generateDigitCaptcha c.Generate() err:%+v", err)
	}
	return
}
