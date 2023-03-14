// 以下是驗證程式碼, 不可使用
// using區段可能與實際使用的不一致, 要注意

using Newtonsoft.Json;

namespace Sheeter
{
    using Data_ = Handmade;
    using Key_ = System.String;
    using Store_ = Dictionary<string, Handmade>;

    /// <summary>
    /// $結構說明
    /// </summary>
    public partial class Handmade
    {
        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Pkey")]
        public int Pkey { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Skey")]
        public string Skey { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data1")]
        public bool Data1 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data2")]
        public bool[] Data2 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data3")]
        public int Data3 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data4")]
        public int[] Data4 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data5")]
        public long Data5 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data6")]
        public long[] Data6 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data7")]
        public float Data7 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data8")]
        public float[] Data8 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data9")]
        public double Data9 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data10")]
        public double[] Data10 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data11")]
        public string Data11 { get; set; }

        /// <summary>
        /// $欄位說明
        /// </summary>
        [JsonProperty("Data12")]
        public string[] Data12 { get; set; }
    }

    /// <summary>
    /// $結構說明
    /// </summary>
    public partial class HandmadeReader : Reader
    {
        /// <summary>
        /// 取得檔名物件
        /// </summary>
        public FileName FileName()
        {
            return new FileName("handmade", ".json");
        }

        /// <summary>
        /// 讀取資料
        /// </summary>
        public string FromData(string data)
        {
            try
            {
                Data = JsonConvert.DeserializeObject<Store_>(data);
            } // try
            catch
            {
                return "from data failed: deserialize failed";
            } // catch

            if (Data == null)
                return "from data failed: deserialize failed";

            return string.Empty;
        }

        /// <summary>
        /// 合併資料
        /// </summary>
        public string MergeData(string data)
        {
            Store_ tmpl;

            try
            {
                tmpl = JsonConvert.DeserializeObject<Store_>(data);
            } // try
            catch
            {
                return "merge data failed: deserialize failed";
            } // catch

            if (tmpl == null)
                return "merge data failed: deserialize failed";

            foreach (var itor in tmpl)
            {
                if (Data.ContainsKey(itor.Key))
                    return "merge data failed: key duplicate";

                Data[itor.Key] = itor.Value;
            } // for

            return string.Empty;
        }

        /// <summary>
        /// 清除資料
        /// </summary>
        public void Clear()
        {
            Data.Clear();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public bool TryGetValue(Key_ key, out Data_ value)
        {
            return Data.TryGetValue(key, out value);
        }

        /// <summary>
        /// 檢查索引是否存在
        /// </summary>
        public bool ContainsKey(Key_ key)
        {
            return Data.ContainsKey(key);
        }

        /// <summary>
        /// 取得迭代器
        /// </summary>
        public IEnumerator<KeyValuePair<Key_, Data_>> GetEnumerator()
        {
            return Data.GetEnumerator();
        }

        /// <summary>
        /// 取得資料
        /// </summary>
        public Data_ this[Key_ key]
        {
            get { return Data[key]; }
        }

        /// <summary>
        /// 取得索引列表
        /// </summary>
        public ICollection<Key_> Keys
        {
            get { return Data.Keys; }
        }

        /// <summary>
        /// 取得資料列表
        /// </summary>
        public ICollection<Data_> Values
        {
            get { return Data.Values; }
        }

        /// <summary>
        /// 取得資料數量
        /// </summary>
        public int Count
        {
            get { return Data.Count; }
        }

        private Store_ Data = new Store_();
    }
}
