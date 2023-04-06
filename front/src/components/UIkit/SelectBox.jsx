import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import { makeStyles } from '@material-ui/styles';

const SelectBox = (props) => {
  const classes = useStyles();

  return (
    <FormControl required={props.required} className={classes.formControl} error={props.error} disabled={props.error}>
      <InputLabel className={classes.selectBoxLabel} error={props.validErr}>{props.label}</InputLabel>
      <Select
        className={classes.selectBox}
        value={props.value}
        onChange={(event) => props.select(event.target.value)}
      >
        {props.options && props.options.map((option) => (
          <MenuItem key={option.id} value={option.id}>
            {option.name}
          </MenuItem>
        ))}
      </Select>
      {props.errorMsg && (
        <FormHelperText error className={classes.selectBoxHelperErrTxt}>{props.errorMsg}</FormHelperText>
      )}
    </FormControl>
  );
};

const useStyles = makeStyles({
  formControl: {
    marginBottom: 16,
    minWidth: 128,
    width: '100%',
  },
  selectBoxLabel: {
    fontSize: '1.3rem',
  },
  selectBox: {
    fontSize: 18,
    paddingTop: 10,
  },
  selectBoxHelperErrTxt: {
    fontSize: '0.9rem',
  }
});

export default SelectBox;
