package query

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/goutils/archivex"
	"github.com/tangx/goutils/filex"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

func init() {
	QueryRouterGroup.Register(&DownloadCertByDomain{})
}

type DownloadCertByDomain struct {
	httpx.MethodGet `path:"/:domain/download"`
	Domain          string `uri:"domain"`
}

func (req *DownloadCertByDomain) Output(c *gin.Context) (interface{}, error) {
	zipfile, err := download(req.Domain)
	if err != nil {
		return httpresponse.RespNotFound(err), nil
	}

	c.FileAttachment(zipfile, zipfile)
	c.Abort()

	return nil, nil
}

func archive(zipfile string, cert legox.Certificate) error {

	keyfile := fmt.Sprintf("%s.key", cert.Domain)
	crtfile := fmt.Sprintf("%s.crt", cert.Domain)

	// 保存文件
	save(keyfile, cert.PrivateKey)
	save(crtfile, cert.Certificate)

	// archive
	err := archivex.ZipFiles(zipfile, keyfile, crtfile)
	if err != nil {
		return err
	}

	// 清理打包后的文件
	os.Remove(keyfile)
	os.Remove(crtfile)
	return nil
}

func save(name string, content string) {
	fobj, _ := os.Create(name)
	defer fobj.Close()
	_, _ = fobj.WriteString(content)
}

func download(domain string) (string, error) {

	// 域名证书是否存在
	cert, ok := utils.GetCert(domain)
	if !ok {
		return "", errors.New("no such domain")
	}

	// 域名证书压缩包是否存在
	zipfile := fmt.Sprintf("%s.zip", domain)
	if filex.Exist(zipfile) {
		// 文件生成不超过 n 天，直接返回
		dayAfter := 80
		fs, _ := os.Stat(zipfile)
		deltaTime := fs.ModTime().AddDate(0, 0, dayAfter).Local()
		now := time.Now().Local()

		if now.Before(deltaTime) {
			return zipfile, nil
		}
	}

	// 创建新的压缩包
	err := archive(zipfile, cert)
	if err != nil {
		return "", err
	}

	return zipfile, nil
}
