import { useCallback, useState, useEffect } from 'react';
import { TextInput, PrimaryButton } from '../components/UIkit';
import { registerEmail } from '../reducks/users/operations';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { getIsLogedIn } from '../reducks/users/selectors';

const EmailRegister = () => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);
  useEffect(() => {
    if (isLogedIn) {
      dispatch(push('/'));
    }
  });

  const [email, setEmail] = useState('');

  const inputEmail = useCallback(
    (event) => {
      setEmail(event.target.value);
    },
    [setEmail]
  );

  return (
    <div className='c-section-container'>
      <h2 className='u-text__headline__k u-text-center'>新規登録</h2>
      <p className='u-text-center'>
        登録用URLをお送りします
        <br />
        メールアドレスを入力して下さい
      </p>
      <div className='module-spacer--medium' />
      <TextInput
        fullWidth={true}
        label={'メールアドレス'}
        multiline={false}
        required={true}
        rows={1}
        value={email}
        type={'email'}
        onChange={inputEmail}
      />
      <div className='module-spacer--medium' />
      <div className='center'>
        <PrimaryButton
          label={'メールアドレスを登録する'}
          onClick={() => dispatch(registerEmail(email))}
          addStyle={{ fontWeight: 'bold' }}
        />
        <div className='module-spacer--medium' />
        <p className='cursor-pointer' onClick={() => dispatch(push('/login'))}>
          アカウントをお持ちの方はこちら
        </p>
      </div>
    </div>
  );
};

export default EmailRegister;
