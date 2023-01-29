import { useCallback, useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { useParams } from 'react-router-dom';
import { push } from 'connected-react-router';
import { makeStyles } from '@material-ui/styles';
import HTMLReactParser from 'html-react-parser';
import { ImageSwiper } from '../components/Products';
import { addProductToCart } from '../reducks/users/operations';
import { showMessageAction } from '../reducks/messages/actions';
import { getWindowSize } from '../reducks/utils/selectors';
import { fetchProduct } from '../reducks/products/operations';
import { constants } from '../utils/constants';
import product_top from '../assets/img/src/nago_product_top.png';
import flow_1 from '../assets/img/svg/nago_flow_1.svg';
import flow_2 from '../assets/img/svg/nago_flow_2.svg';
import flow_3 from '../assets/img/svg/nago_flow_3.svg';
import flow_4 from '../assets/img/svg/nago_flow_4.svg';
import flow_5 from '../assets/img/svg/nago_flow_5.svg';
import flow_6 from '../assets/img/svg/nago_flow_6.svg';
import { theme } from '../assets/theme';

const ProductDetail = () => {
  const classes    = useStyles();
  const dispatch   = useDispatch();
  const selector   = useSelector((state) => state);
  const windowSize = getWindowSize(selector);

  const { productId } = useParams();
  if (!productId) {
    dispatch(showMessageAction('error', '無効なURLです'));
    dispatch(push('/'));
  } else {
    useEffect(() => {
      dispatch(fetchProduct(productId));
    }, []);
  }

  //const path = selector.router.location.pathname;

  const returnCodeToBr = (text) => {
    if (text === '') {
      return text;
    } else {
      return HTMLReactParser(text.replace(/\r?\n/g, '<br/>'));
    }
  };

  const renderDescription = () => {
    return (
      <div className={classes.secondSectionSubDescription}>
        <p>「一人暮らしの高齢者をITの力で見守りたい」</p>
        <p>祖父の転倒死をきっかけにそう思うようになりました。</p>
        <p>本サービスはそんな思いから生まれた</p>
        <p>IoT機器とLINEで実現する</p>
        <p>一人暮らし高齢者向け見守り通知サービスです。</p>
      </div>
    );
  };

  //const addProduct = useCallback();
  //(selectedSize) => {
  //  //const timestamp = FirebaseTimestamp.now();
  //  dispatch(
  //    addProductToCart({
  //      //added_at: timestamp,
  //      description: product.description,
  //      gender: product.gender,
  //      images: product.images,
  //      name: product.name,
  //      price: product.price,
  //      productId: product.id,
  //      quantity: 1,
  //    })
  //  );
  //},
  //[dispatch, product]
  const product = constants.product;

  return (
    <section
      className={classes.sectionWrapperAdjuster}
    >
      {/* セクション1 */}
      <div className='topSection'>
        {product && (
          <div className='p-grid__row'>
            <div className={classes.sliderBox}>
              <ImageSwiper images={product.images} />
            </div>
            <div className={classes.detail}>
              <p>{product.subTitle}</p>
              <h2
                className={'u-text__headline' + ' ' + `${classes.productTitle}`}
              >
                {product.name}
              </h2>
              <p className={classes.price}>¥{product.price.toLocaleString()}</p>
              <div className='module-spacer--small' />
              <div className='module-spacer--small' />
              <p>{returnCodeToBr(product.description)}</p>
            </div>
          </div>
        )}
      </div>
      {/* セクション2 */}
      <div className={`${classes.colorSection} ${classes.secondSection}`}>
        <div className={classes.secondSectionSubWrapper}>
          <div className={classes.secondSectionSub}>
            <p className={classes.secondSectionSubSubTitle}>
              あなたの大切な人は&quot;今&quot;どうしてますか?
            </p>
            <p className={classes.secondSectionSubTitle}>
              IoT見守りサービス
              <br />
              「nago」
            </p>
            {windowSize.width >= theme.breakpoints.values.sm &&
              renderDescription()}
          </div>
          <div className={classes.secondSectionSub}>
            <img src={product_top} />
          </div>
          {windowSize.width < theme.breakpoints.values.sm &&
            renderDescription()}
        </div>
      </div>
      {/* セクション3 */}
      <div className={classes.thirdSection}>
        <p className={classes.thirdSectionTitle}>nagoで見守りを始めるまで</p>
        <div className={classes.thirdSectionFlow}>
          <div>
            <img src={flow_1} />
            <p>1.こちらのサイトから購入します</p>
          </div>
          <div>
            <img src={flow_2} />
            <p>2.ご自宅に商品が届きます</p>
          </div>
          <div>
            <img src={flow_3} />
            <p>3.見守りたい方が普段良く通る場所に機器を設置します</p>
          </div>
          <div>
            <img src={flow_4} />
            <p>4.LINEの登録を行います</p>
          </div>
          <div>
            <img src={flow_5} />
            <p>5.通知設定を行います</p>
          </div>
          <div>
            <img src={flow_6} />
            <p>
              6.設定した時間を超えても通過していない場合LINEに通知します。通過するたびに通知することも可能です
            </p>
          </div>
        </div>
      </div>
    </section>
  );
};

const useStyles = makeStyles((theme) => ({
  sectionWrapperAdjuster: {
    marginTop: '50px',
  },
  colorSection: {
    color: theme.palette.primary.light,
    backgroundColor: '#4dd0e1',
  },
  // セクション1
  sliderBox: {
    margin: '0 auto 24px auto',
    [theme.breakpoints.down('sm')]: {
      width: 320,
    },
    [theme.breakpoints.up('sm')]: {
      width: 400,
    },
  },
  productTitle: {
    fontWeight: 'bold',
  },
  detail: {
    textAlign: 'left',
    [theme.breakpoints.down('sm')]: {
      margin: '0 auto 16px auto',
      height: 'auto',
      width: 320,
    },
    [theme.breakpoints.up('sm')]: {
      margin: '0 auto',
      height: 'auto',
      width: 400,
    },
  },
  price: {
    fontSize: 36,
  },
  // セクション2
  secondSection: {
    height: 'calc(100vh - 72px)',
    display: 'flex',
  },
  secondSectionSubWrapper: {
    display: 'flex',
    width: '100%',
    maxWidth: theme.size.window.contentMaxWidth,
    margin: '0 auto',
    flexDirection: 'column',
    [theme.breakpoints.down('xs')]: {
      // SP(〜600px)
      marginTop: '32px',
    },
    [theme.breakpoints.up('sm')]: {
      // タブレット以上(600px〜)では横向き
      flexDirection: 'row',
    },
  },
  secondSectionSub: {
    flex: 1,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    flexFlow: 'column',
    textAlign: 'center',
    '& > img': {
      maxHeight: '100%',
    },
    [theme.breakpoints.down('xs')]: {
      // SP で img がはみ出るため
      maxHeight: 'calc(100vh / 3)',
    },
  },
  secondSectionSubSubTitle: {
    width: '100%',
    fontSize: '18px',
    [theme.breakpoints.up('md')]: {
      // PC(960px〜)
      fontSize: '24px',
    },
  },
  secondSectionSubTitle: {
    width: '100%',
    fontWeight: 'bold',
    fontSize: '30px',
    [theme.breakpoints.down('xs')]: {
      // SP
      marginTop: 0,
    },
    [theme.breakpoints.up('sm')]: {
      // タブレット
      fontSize: '32px',
    },
    [theme.breakpoints.up('md')]: {
      // PC
      fontSize: '36px',
    },
  },
  secondSectionSubDescription: {
    flexDirection: 'column',
    marginTop: '32px',
    textAlign: 'center',
    [theme.breakpoints.down('xs')]: {
      // SP で img がはみ出るため
      maxHeight: 'calc(100vh / 3)',
      marginTop: 0,
      paddingBottom: '20px',
    },
    '& > p': {
      marginTop: 2,
      fontSize: '14px',
      [theme.breakpoints.up('md')]: {
        // PC
        fontSize: '20px',
      },
    },
  },
  // セクション3
  thirdSection: {
    padding: '40px 20px 0',
    [theme.breakpoints.up('md')]: {
      // PC
      padding: '60px 40px',
    },
  },
  thirdSectionTitle: {
    textAlign: 'center',
    fontWeight: 'bold',
    marginBottom: '32px',
    [theme.breakpoints.down('xs')]: {
      // SP
      fontSize: '20px',
    },
    [theme.breakpoints.up('sm')]: {
      // タブレット
      fontSize: '24px',
    },
    [theme.breakpoints.up('md')]: {
      // PC
      fontSize: '28px',
    },
  },
  thirdSectionFlow: {
    display: 'flex',
    flexWrap: 'wrap',
    flexDirection: 'row',
    '& > div': {
      display: 'flex',
      flexFlow: 'column',
      alignSelf: 'baseline',
      padding: '16px 8px',
      [theme.breakpoints.down('xs')]: {
        // SP
        width: '100%',
      },
      [theme.breakpoints.up('sm')]: {
        // タブレット
        width: 'calc(100% / 2)',
      },
      [theme.breakpoints.up('md')]: {
        // PC
        width: 'calc(100% / 3)',
      },
      '& > img': {
        flexBasis: '90%',
      },
      '& > p': {
        flexBasis: '10%',
        [theme.breakpoints.down('xs')]: {
          // SP
          fontSize: '16px',
        },
        [theme.breakpoints.up('sm')]: {
          // タブレット
          fontSize: '18px',
        },
        [theme.breakpoints.up('md')]: {
          // PC
          fontSize: '18px',
        },
      },
    },
  },
}));

export default ProductDetail;
