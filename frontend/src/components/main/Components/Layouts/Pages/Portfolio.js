import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";
import Item from "./PortfolioItems";
import { MDBContainer, MDBRow, MDBCol } from "mdbreact";

import "./page.css";
import { getPortfolios } from "../../../../../utils/backend";

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <Typography
      component="div"
      role="tabpanel"
      hidden={value !== index}
      id={`scrollable-auto-tabpanel-${index}`}
      aria-labelledby={`scrollable-auto-tab-${index}`}
      {...other}
    >
      <Box p={3}>{children}</Box>
    </Typography>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.any.isRequired,
  value: PropTypes.any.isRequired
};

function a11yProps(index) {
  return {
    id: `scrollable-auto-tab-${index}`,
    "aria-controls": `scrollable-auto-tabpanel-${index}`
  };
}

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    width: "100%",
    backgroundColor: theme.palette.background.paper
  }
}));

export default function Portfolio() {
  const classes = useStyles();
  const [value, setValue] = React.useState(0);
  const [portfolio, setPortfolio] = React.useState([]);
  const [sectors, setSectors] = React.useState([]);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  const GetPortfolio = () => {
    getPortfolios(null, (err, res) => {
      if (err) {
        console.log(err);
      } else {
        setPortfolio(res.toObject().portfoliosList);

        setSectors(res.toObject().portfoliosList.map(data => data.sector));
      }
    });
  };

  useEffect(() => {
    GetPortfolio();
  }, []);

  return (
    <MDBContainer id="portfolio">
      <div
        className={classes.root}
        style={{ textAlign: "center", margin: "0" }}
      >
        <h3>Portfolio</h3>
        <div>
          <AppBar
            position="static"
            style={{
              color: "#0dc835",
              width: "60%",
              margin: "auto",
              background: "transparent",
              marginTop: "3%",
              boxShadow: "none",
              zIndex: "0"
            }}
          >
            <Tabs
              value={value}
              onChange={handleChange}
              indicatorColor="primary"
              textColor="primary"
              variant="scrollable"
              scrollButtons="auto"
              aria-label="scrollable auto tabs example"
              id="textcolor"
            >
              <Tab label="All" {...a11yProps(0)} />
              {sectors.map((s, i) => (
                <Tab key={i} label={s} {...a11yProps(i + 1)} />
              ))}
            </Tabs>
          </AppBar>
        </div>

        <TabPanel value={value} index={0}>
          <MDBContainer>
            <MDBRow>
              {portfolio.map((item, i) => (
                <MDBCol key={i} md={4} style={{ minHeight: "240px" }}>
                  <Item item={item} />
                </MDBCol>
              ))}
            </MDBRow>
          </MDBContainer>
        </TabPanel>
        {sectors.map((s, i) => (
          <TabPanel key={i} value={value} index={i + 1}>
            <MDBRow>
              {portfolio
                .filter(item => item.sector === s)
                .map((item, j) => (
                  <MDBCol key={j} md={4} style={{ minHeight: "240px" }}>
                    <Item item={item} />
                  </MDBCol>
                ))}
            </MDBRow>
          </TabPanel>
        ))}
      </div>
    </MDBContainer>
  );
}
