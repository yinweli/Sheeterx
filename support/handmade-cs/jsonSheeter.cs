// 以下是驗證程式碼, 不可使用
// using區段可能與實際使用的不一致, 要注意

namespace Sheeter
{
    /// <summary>
    /// 表格資料
    /// </summary>
    public partial class Sheeter
    {
        public Sheeter(Loader loader)
        {
            this.loader = loader;
            this.reader = new List<Reader>() { this.Handmade, };
        }

        /// <summary>
        /// 讀取資料處理
        /// </summary>
        public bool FromData()
        {
            if (loader == null)
                return false;

            var result = true;

            foreach (var itor in reader)
            {
                var filename = itor.FileName();
                var data = loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var error = itor.FromData(data);

                if (error.Length != 0)
                {
                    result = false;
                    loader.Error(filename.File, error);
                } // if
            } // for

            return result;
        }

        /// <summary>
        /// 合併資料處理
        /// </summary>
        public bool MergeData()
        {
            if (loader == null)
                return false;

            var result = true;

            foreach (var itor in reader)
            {
                var filename = itor.FileName();
                var data = loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var error = itor.MergeData(data);

                if (error.Length != 0)
                {
                    result = false;
                    loader.Error(filename.File, error);
                } // if
            } // for

            return result;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            foreach (var itor in reader)
                itor.Clear();
        }

        /// <summary>
        /// 裝載器物件
        /// </summary>
        private readonly Loader loader;

        /// <summary>
        /// 讀取器列表
        /// </summary>
        private readonly List<Reader> reader;

        /// <summary>
        /// $表格說明
        /// </summary>
        public readonly HandmadeReader Handmade = new HandmadeReader();
    }

    /// <summary>
    /// 裝載器介面
    /// </summary>
    public interface Loader
    {
        /// <summary>
        /// 讀取檔案
        /// </summary>
        public string Load(FileName filename);

        /// <summary>
        /// 錯誤處理
        /// </summary>
        public void Error(string name, string message);
    }

    /// <summary>
    /// 讀取器介面
    /// </summary>
    public interface Reader
    {
        /// <summary>
        /// 取得檔名物件
        /// </summary>
        public FileName FileName();

        /// <summary>
        /// 讀取資料
        /// </summary>
        public string FromData(string data);

        /// <summary>
        /// 合併資料
        /// </summary>
        public string MergeData(string data);

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear();
    }

    /// <summary>
    /// 檔名資料
    /// </summary>
    public class FileName
    {
        public FileName(string name, string ext)
        {
            this.name = name;
            this.ext = ext;
        }

        /// <summary>
        /// 取得名稱
        /// </summary>
        public string Name
        {
            get { return name; }
        }

        /// <summary>
        /// 取得副檔名
        /// </summary>
        public string Ext
        {
            get { return ext; }
        }

        /// <summary>
        /// 取得完整檔名
        /// </summary>
        public string File
        {
            get { return name + ext; }
        }

        /// <summary>
        /// 名稱
        /// </summary>
        private readonly string name;

        /// <summary>
        /// 副檔名
        /// </summary>
        private readonly string ext;
    }
}
