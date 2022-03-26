import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import PersonAddIcon from '@material-ui/icons/PersonAdd';
import HistoryIcon from '@material-ui/icons/History';
import PersonIcon from '@material-ui/icons/Person';
import EmojiPeopleIcon from '@material-ui/icons/EmojiPeople';

const Config = {
  apiUrl: process.env.REACT_APP_API_URL,
  menus: {
    noAuthenticated: [
      {
        label: 'ログイン',
        icon: <ExitToAppIcon />,
        value: '/login',
      },
      {
        label: '新規登録',
        icon: <PersonAddIcon />,
        value: '/email_register',
      },
    ],
    authenticated: [
      {
        label: '注文履歴',
        icon: <HistoryIcon />,
        value: '/order/history',
      },
      {
        label: 'プロフィール',
        icon: <PersonIcon />,
        value: '/user/mypage',
      },
      {
        label: 'ログアウト',
        icon: <EmojiPeopleIcon />,
        value: 'logout',
      },
    ],
  },
};

export default Config;
