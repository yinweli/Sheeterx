// 以下是模板驗證程式碼

package sheeter

// NewSheeter 建立表格資料
func NewSheeter(loader Loader) *Sheeter {
	sheeter := &Sheeter{}
	sheeter.loader = loader
	sheeter.reader = []Reader{
		&sheeter.Handmade,
	}
	return sheeter
}

// Sheeter 表格資料
type Sheeter struct {
	loader Loader   // 裝載器物件
	reader []Reader // 讀取器列表

	Handmade HandmadeReader // $表格說明
}

// FromData 讀取資料處理
func (this *Sheeter) FromData() bool {
	if this.loader == nil {
		return false
	} // if

	result := true

	for _, itor := range this.reader {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		} // if

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		} // if
	} // for

	return result
}

// MergeData 合併資料處理
func (this *Sheeter) MergeData() bool {
	if this.loader == nil {
		return false
	} // if

	result := true

	for _, itor := range this.reader {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		} // if

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		} // if
	} // for

	return result
}

// Clear 清除資料
func (this *Sheeter) Clear() {
	for _, itor := range this.reader {
		itor.Clear()
	} // for
}

// Loader 裝載器介面
type Loader interface {
	// Load 讀取檔案
	Load(filename FileName) []byte

	// Error 錯誤處理
	Error(name string, err error)
}

// Reader 讀取器介面
type Reader interface {
	// FileName 取得檔名物件
	FileName() FileName

	// FromData 讀取資料
	FromData(data []byte) error

	// MergeData 合併資料
	MergeData(data []byte) error

	// Clear 清除資料
	Clear()
}

// NewFileName 建立檔名資料
func NewFileName(name, ext string) FileName {
	return FileName{
		name: name,
		ext:  ext,
	}
}

// FileName 檔名資料
type FileName struct {
	name string // 名稱
	ext  string // 副檔名
}

// Name 取得名稱
func (this FileName) Name() string {
	return this.name
}

// Ext 取得副檔名
func (this FileName) Ext() string {
	return this.ext
}

// File 取得完整檔名
func (this FileName) File() string {
	return this.name + this.ext
}
