import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
import { makeStyles } from '@material-ui/core/styles';
import { useDispatch, useSelector } from 'react-redux';
import {
  getMessageState,
  getMessageSeverity,
  getMessageText,
} from '../../reducks/messages/selectors';
import { hideMessageAction } from '../../reducks/messages/actions';

const useStyles = makeStyles(() => ({
  root: {
    position: 'fixed',
    top: '90px',
    opacity: 0.9,
  },
  multiLine: {
    // 複数メッセージの場合改行
    whiteSpace: 'pre-wrap',
  },
}));

const Message = ({ children }) => {
  const dispatch = useDispatch();

  const classes = useStyles();

  const selector = useSelector((state) => state);
  const isShowingMessage = getMessageState(selector);
  const showingSeverity = getMessageSeverity(selector);
  let showingText = getMessageText(selector);
  if (Array.isArray(showingText)) {
    showingText = showingText.join('\n');
  }

  const handleClose = (event, reason) => {
    // メッセージコンポーネント外をクリックした場合は閉じない
    if (reason === 'clickaway') return;
    // エラーの場合フラッシュメッセージではなく表示し続ける
    if (showingSeverity === 'error' && reason === 'timeout') return;
    dispatch(hideMessageAction());
  };

  return (
    <>
      {isShowingMessage && (
        <Snackbar
          anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
          open={isShowingMessage}
          className={classes.root}
          onClose={handleClose}
          autoHideDuration={5000}
        >
          <Alert
            className={classes.multiLine}
            severity={showingSeverity}
            onClose={handleClose}
          >
            {showingText}
          </Alert>
        </Snackbar>
      )}
      {children}
    </>
  );
};

export default Message;
