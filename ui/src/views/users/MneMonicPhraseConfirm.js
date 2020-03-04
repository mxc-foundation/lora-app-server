import React, { useState } from "react";
import { Row, Col, Button } from 'reactstrap';
import classNames from "classnames";

import i18n, { packageNS } from '../../i18n';


const Phrase = ({ phrase, isSelected, select }) => {
    return <React.Fragment>
        <Button color={isSelected ? "primary" : "secondary"} outline className={classNames("btn-rounded", "m-1", { "bg-white": isSelected, "text-primary": isSelected })} onClick={select}>
            {phrase}
        </Button>
    </React.Fragment>
}

const MneMonicPhraseConfirm = ({ title, phrase, next, back }) => {

    const [selectedPhrase, setSelectedPhrase] = useState([]);

    const onSelect = (phrase, isSelected) => {
        let phrases = [...selectedPhrase];
        if (isSelected) {
            phrases = phrases.filter(p => p !== phrase);
        } else {
            phrases.push(phrase);
        }
        setSelectedPhrase(phrases);
    }

    return <React.Fragment>
        <Row className="text-center">
            <Col className="mb-0">
                <h5>{title}</h5>

                <Row className="mt-3 text-left">
                    <Col className="mb-0">
                        <div className="bg-light p-3">
                            {selectedPhrase.map((word, idx) => (
                                <Phrase key={idx} phrase={word} isSelected={true} select={() => { onSelect(word, true) }} />
                            ))}
                        </div>
                    </Col>
                </Row>

                <Row className="mt-3 text-left">
                    <Col className="mb-0">
                        <div>
                            {(phrase || []).map((word, idx) => {
                                return selectedPhrase.indexOf(word) === -1 ? <Phrase key={idx} phrase={word} isSelected={false} select={() => { onSelect(word, false) }} />: null
                            })}
                        </div>
                    </Col>
                </Row>

                <Row className="mt-2 text-left">
                    <Col className="mb-0">
                        <Button color="primary" className="btn-block" onClick={() => next(selectedPhrase)}
                            disabled={!selectedPhrase.length}>{i18n.t(`${packageNS}:menu.menmonic_phrase.confirm_button`)}</Button>
                    </Col>
                    <Col className="mb-0">
                        <Button color="primary" outline className="btn-block" onClick={back}>{i18n.t(`${packageNS}:menu.menmonic_phrase.back_button`)}</Button>
                    </Col>
                </Row>

            </Col>
        </Row>
    </React.Fragment>
}

export default MneMonicPhraseConfirm;
