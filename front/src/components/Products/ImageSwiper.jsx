import { useState } from 'react';
import Swiper from 'react-id-swiper';
import NoImage from '../../assets/img/src/no_image.png';
import 'swiper/css/swiper.css';
import { TransformWrapper, TransformComponent } from 'react-zoom-pan-pinch';

const ImageSwiper = (props) => {
  const [params] = useState({
    grabCursor: true,
    centeredSlides: true,
    spaceBetween: 30,
    loop: true,
    pagination: {
      el: '.swiper-pagination',
      type: 'fraction',
      clidkable: true,
    },
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
  });

  const images = props.images;

  return (
    <Swiper {...params}>
      {images.length === 0 ? (
        <div className='p-media__thumb'>
          <img src={NoImage} alt='no image' />
        </div>
      ) : (
        images.map((image) => (
          <div className='p-media__thumb' key={image.id}>
            <TransformWrapper pinch={{ step: 1 }} doubleClick={{ step: 0.5 }}>
              <TransformComponent>
                <img src={image.path} alt='商品画像' />
              </TransformComponent>
            </TransformWrapper>
          </div>
        ))
      )}
    </Swiper>
  );
};

export default ImageSwiper;
