import { useSelector, useDispatch } from 'react-redux';
import { push } from 'connected-react-router';
import { makeStyles } from '@material-ui/styles';
import product_top from '../assets/img/src/nago_product_top.png';
import flow_1 from '../assets/img/svg/nago_flow_1.svg';
import flow_2 from '../assets/img/svg/nago_flow_2.svg';
import flow_3 from '../assets/img/svg/nago_flow_3.svg';
import flow_4 from '../assets/img/svg/nago_flow_4.svg';
import flow_5 from '../assets/img/svg/nago_flow_5.svg';
import flow_6 from '../assets/img/svg/nago_flow_6.svg';
import { theme } from '../assets/theme';
import { getWindowSize } from '../reducks/utils/selectors';
import { TextButton } from '../components/UIkit';

const Home = () => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const windowSize = getWindowSize(selector);

  const renderDescription = () => {
    return (
      <div className={classes.topSectionSubDescription}>
        <p>「一人暮らしの高齢者をITの力で見守りたい」</p>
        <p>祖父の転倒死をきっかけにそう思うようになりました。</p>
        <p>本サービスはそんな思いから生まれた</p>
        <p>IoT機器とLINEで実現する</p>
        <p>一人暮らし高齢者向け見守り通知サービスです。</p>
      </div>
    );
  };

  return (
    <section>
      <div className={`${classes.colorSection} ${classes.topSection}`}>
        <div className={classes.topSectionSubWrapper}>
          <div className={classes.topSectionSub}>
            <p className={classes.topSectionSubSubTitle}>
              あなたの大切な人は&quot;今&quot;どうしてますか?
            </p>
            <p className={classes.topSectionSubTitle}>
              IoT見守りサービス
              <br />
              「nago」
            </p>
            {windowSize.width >= theme.breakpoints.values.sm &&
              renderDescription()}
          </div>
          <div className={classes.topSectionSub}>
            <img src={product_top} />
          </div>
          {windowSize.width < theme.breakpoints.values.sm &&
            renderDescription()}
        </div>
      </div>
      <div className={classes.secondSection}>
        <p className={classes.secondSectionTitle}>nagoで見守りを始めるまで</p>
        <div className={classes.secondSectionFlow}>
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
      <div className={`${classes.colorSection} ${classes.bottomSection}`}>
        <TextButton
          label={'商品詳細・購入ページはこちら'}
          onClick={() => dispatch(push('/product'))}
        />
      </div>
    </section>
  );
};

// スタイル
const useStyles = makeStyles(() => ({
  colorSection: {
    color: theme.palette.primary.light,
    backgroundColor: '#4dd0e1',
  },
  // セクション1
  topSection: {
    height: 'calc(100vh - 72px)',
    display: 'flex',
  },
  topSectionSubWrapper: {
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
  topSectionSub: {
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
  topSectionSubSubTitle: {
    width: '100%',
    fontSize: '18px',
    [theme.breakpoints.up('md')]: {
      // PC(960px〜)
      fontSize: '24px',
    },
  },
  topSectionSubTitle: {
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
  topSectionSubDescription: {
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
  // セクション2
  secondSection: {
    padding: '40px 20px 0',
    [theme.breakpoints.up('md')]: {
      // PC
      padding: '60px 40px',
    },
  },
  secondSectionTitle: {
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
  secondSectionFlow: {
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
  // セクション3
  bottomSection: {
    position: 'fixed',
    bottom: 0,
    left: 0,
    right: 0,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: '64px',
    [theme.breakpoints.down('xs')]: {
      // SP
      height: '56px',
    },
  },
}));

export default Home;
