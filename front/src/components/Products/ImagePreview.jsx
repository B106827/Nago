const ImagePreview = (props) => {
  return (
    // 画像を1:1で切り取るCSS
    <div className='p-meida__thumb' onClick={() => props.delete(props.id)}>
      <img alt='プレビュー画像' src={props.path} />
    </div>
  );
};

export default ImagePreview;
