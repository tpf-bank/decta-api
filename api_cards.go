/*
Decta API

<h2>Introduction</h2><p>This is a documented guide that    describes the process of DECTA card and token operations for both private and business clients. As a new or already    experienced API user you will find in this documentation all necessary and required information to start using DECTA    API and improve your experience with DECTA.</p><h2>Authentication and Authorisation</h2><p>Each request to DECTA API must be signed with a certificate that allows DECTA to identify API clients. This    guarantees the authenticity of the request received from the client.</p><h3>Authentication</h3><p>The authentication process is executed only once for each client.<br>    On successful authentication DECTA assigns a certificate to the client which must be used in every API    request in the future.</p><p>    To access a detailed description of the authentication process, choose <b>Service</b> specification at the top right    corner or click <a        href=\"/?urls.primaryName=SERVICE\" target=\"_blank\">here</a>.</p><h3>Authorisation</h3><p>    Each request must be signed using Jose specifications in particular - <a        href=\"https://tools.ietf.org/html/rfc7515\" target=\"_blank\">RFC7515</a>. It is vital for users to follow the    provided standard and use <b>Base64url    Encoding without Padding</b> for token encoding.</p><h3>Request with filters</h3><p>    Several requests support data filters. Filter syntax can be found below -<p>    <code>        &lt;requestFilter&gt; ::= &lt;filterName&gt;=[eq|ne|gt|ge|lt|le]:&lt;filterValue&gt;    </code></p><ul>Where:    <li>eq - equals</li>    <li>ne - not equals</li>    <li>gt - greater than</li>    <li>ge - greater or equals than</li>    <li>lt - less than</li>    <li>le - less or equals than</li></ul><ul>Example:    <li>cardState: \"eq:BLOCKED_BY_HOLDER\"</li></ul><h2>Changelog</h2>Information about the latest DECTA API changes and improvements can be found <a href=\"/changelog.html\" target=\"_blank\">here.</a>.<h4>Request examples</h4><p>    We recommend users follow the instructions in the next sections below and create the application step by step.</p><p>    To understand the process of DECTA API requests better, we have prepared two basic API calls as examples.</p><p>    <b>POST /v1/api/cards/order</b> request</p><p>For example:</p><ul>    <li>        <b>uri</b> - <br>/v1/api/cards/order    </li>    <li>        <b>body</b> - <br>        {\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>        308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100B73748751B973456931B380163CE168BF884B26B1D27FA6978ACCEF829DF2DBA5D2845ABE997E82531FFD1C22D5E9F1ED6672F6D75D6572B688AE24AB9ED5138D5B5AA484FA04CEC393DB84C833BDDF7507BA1F5DC59030BE6E0FAC16613FADF3D06123621C11FB8F64FE777B9D0308DCEDCBAF78EB76FDAA4151C96E48B68507FA254AEFE44F0C0D103CAFE2A4ADB676E2B4FE26DB4AD671B8F02E74FE4B7E36156F44AA7CCE39DB0879926CEE2EF43502C7F99A1A3DDE72ECAC39E83F7B6819F2C1B8606537D6A04A1D6CC1FF0044F00360EA2F899BD9D9C6D748B20E2EB17CF908D6117B6B5B97490F5D68DC2D039B11EB2D920B8CEE9EFF10D0408A69971020301000102820100584078D67003D5621E59EE103D52F7E9543C16F521863048BDA7FFC9E9E166D56E2A07E5570FA4F2C5B2C40714738F91FB1498F3D0DEFDFE1ACD4B53535BCCD3E39D2DF8C6E2202B692721AF39478D13A3E0E992D420CC26DF7F5F49E9319696117EEB26C7FB8E9C398923A5B80B6057EE5CC4729A7C2194DF948BC0E3358378F1E3381BB62FE98C9372BB6EBBC24E8670E80EF991173D6D6347ADF1F8C9B67920D88072978D025F16F50273EDA0F0AA6FABC6F9687259A21142FD1E5C3DE99EE949DF1E81C7E7CDA4DD806A5CE2105DDC0CF7A2274913971C234D04D7C577C76F2FAE75379A0582FFA0333A01A7BE29ECA203AD7E3E71271EDEFDC0069F684102818100E7C26736089B4FB37BE4873EB2A726BBCD9CA3A3F3EB6BAD1433127614ABC7E6763D118A4A8DB45BBE31A67A853A44BA21EDF50533ACD348A0A9C1F648EE510426ECC845C01185C6793E9D4866281B9C01AE8BA96D0897BCBB0F638696CD6935794506AA3EC14BAB1A7F68DDF70DA21C70249FAD3FE88AB4C6708691327E4BC302818100CA611446E3A9A1997D175BDEADDE02B5C09E3E76024741F2DAAEF8689CF5DCA1C5FA9AE8A6F2F09E837E9D181C5607374DFBAA53B02C3908DCFA1817DEF1797F64743218B1616B55125E618B45F11003BA94D89E357358BD54A1DA324D893FDD9C0B30428A770F410B061C011CDB541E344B4C01B17CB615A1DCAD1F5EEB96BB02818100E40750C9C75A18E73E053254AC2EEE5B6608B2B18433A4341D65CACA47B864ED0A7537A6DB87E56747114EFDC9CBF507368F0CBF5B82B638056C419D73509881FF528612AAD212CF9F47CE3507DE7A9BDAC3C442A5370924F6E0434A8F61F81C56FF65796859837C0C8C43BFF16E868C78827061643A070FAB17D82F508117450281804FE4A32C99138E4819A9EF0AA978CB7914E163A7129F2ED9C09AF255DA20F548A7EF96D7E190668D2D3BFEA856076031E50744E664D6106DFF4E7BD4709EC368173007D6D7AFADBF97D0CA9140BB39A73F312392D16707D13667EECB8CF071D5FA94302914A08BD5119507D9289B2D49FF3AFA7670AADAF70F3F1ED9138FCDA102818100E69108E3E83702F3258503C2FD0E9A34E76750CCE2A20AFDD58F3939351979ECC66373A2151AD313710B2EAFA192DF9F364791FA8473CE1CAE434A044AA8F557F99D756DF4373B5A3F2202ED299979F3C5B4AC1B93ABDB4157ED3901B723B671749DE9279D95759F122400DD5BF3B28F6A8A37383ADFAD9E5C5D49192B20EB43    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>        WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br> {\"x5t#S256\":\"WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b>-<br>        /v1/api/cards/order{\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>        eyJhbGciOiJSUzI1NiIsIng1dCNTMjU2IjoiV09RRUV5cmNIYVZvYkhNUE9SbXVNZTlpQUM1Z1NJMFRwWHlJZkExbEtJMCJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>        eyJhbGciOiJSUzI1NiIsIng1dCNTMjU2IjoiV09RRUV5cmNIYVZvYkhNUE9SbXVNZTlpQUM1Z1NJMFRwWHlJZkExbEtJMCJ9.L3YxL2FwaS9jYXJkcy9vcmRlcnsiY2FyZCI6eyJhY2NvdW50T3duZXJSZWxhdGlvbiI6Ik9XTkVSIiwiY3VycmVuY2llcyI6WyJFVVIiXSwiZGVsaXZlcnlBZGRyZXNzIjp7ImNpdHkiOiJNYWRvbmEiLCJjb3VudHJ5IjoiTFZBIiwibmFtZSI6IlRlc3QiLCJwaG9uZSI6OTExMywic2hpcG1lbnQiOiJTVEFOREFSRCIsInN0cmVldCI6IjYyIE1hZG9uYSBzdHJlZXQiLCJzdXJuYW1lIjoiVGVzdFMiLCJ6aXBDb2RlIjoiTFYtMTAwMCJ9LCJob2xkZXIiOnsiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2EiLCJjb3VudHJ5IjoiTFZBIiwic3RyZWV0IjoiNyBTdHJlZXQiLCJ6aXBDb2RlIjoiTFYwMDAwIn0sImRvY3VtZW50Ijp7ImJpcnRoRGF0ZSI6IjE5NTAtMTItMzEiLCJjb3VudHJ5Q29kZSI6IkVVUiIsImV4cGlyeURhdGUiOiIyMDU2LTEyLTMxIiwiaXNzdWluZ0RhdGUiOiIyMDE2LTEyLTMxIiwibnVtYmVyIjoiVFRIb2xkZXIiLCJzdWJ0eXBlIjoiUkVTSURFTlRfUEVSTUlUX0lEIiwidHlwZSI6IkRSSVZJTkdfTElDRU5TRSJ9LCJlbWFpbCI6Im9ob2xkZXJAZGVjdGEuY29tIiwibGFuZ3VhZ2UiOiJFTiIsIm1haWRlbk5hbWUiOiJUZXN0TSIsIm1vYmlsZVBob25lIjoiMTIzNDU2NzgiLCJuYW1lIjoiVGVzdFMiLCJyb2xlIjoiT1dORVIiLCJzdXJuYW1lIjoiVGVzdFMiLCJ1c2VyRGVmaW5lZEZpZWxkMSI6IjIifSwicGFzc3BocmFzZSI6InRlc3QycGFzcyIsInByaW9yaXR5IjoiVVJHRU5UIiwicHJvZHVjdENvZGUiOiI3MDEiLCJzdXBwbGVtZW50YXJ5IjoiZmFsc2UifSwiZXh0ZXJuYWxJZCI6IlByaXZhdGVfTmV3IiwicHJpdmF0ZUNsaWVudCI6eyJjb21tZW50IjoiQWdyZWVtZW50IG51bS4gMTIzMTQzIiwiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2FDdXJyZW50IiwiY291bnRyeSI6IkxWQSIsIm1vbnRocyI6NCwic3RyZWV0IjoiNiBzdHJlZXQgc3RyZWV0IiwiemlwQ29kZSI6IldTY3VyMTAwOSJ9LCJkb2N1bWVudCI6eyJiaXJ0aERhdGUiOiIxOTUwLTEyLTMxIiwiY291bnRyeUNvZGUiOiJMVkEiLCJleHBpcnlEYXRlIjoiMjA5OS0xMi0zMSIsImlzc3VpbmdEYXRlIjoiMTk3MS0xMi0zMSIsIm51bWJlciI6IjEyMzQiLCJzdWJ0eXBlIjoiTkFUSU9OQUxfSUQiLCJ0eXBlIjoiSURfQ0FSRCJ9LCJlbWFpbCI6ImVtYWlsMTMwIiwibGFuZ3VhZ2UiOiJlbiIsIm1haWRlbk5hbWUiOiJUZXN0IiwibW9iaWxlUGhvbmUiOiIzNzEyODg2NjY5NCIsIm5hbWUiOiJUZXN0IiwicGVwIjoiTk8iLCJzdXJuYW1lIjoiUHJpdmF0ZUJhc2VOZXciLCJ1c2VyRGVmaW5lZEZpZWxkMSI6IjIwMDA2NjYifX0.Ks767mp1eA-e1y27GpSmLBhKLud8e96bgYPUpIIlYq5ybeKZ2-WwJsExIIcrnIixaev0ibZ3RrHueKl_uhdF4bVNLiCo8EU2sDnK16e5Z9NL1lS4KZI9WUZHhCIzYJnljdBbHiISSknjdsw3WcvGk6bAYj93U0WplvFxa5lk6Gn2T0lkzujUhJXSHkbuZ541rhwwyztndFbUW07vuTCQ-z_gZuRz8dRbvSnAl0aZdeNWBj0PH3r1GOMG_SR_hpyrf0oEItFGj1ASdOfQjNH78dcMIzle9qh4NzLX8voFTs3u58ktfdK1HhcScumtqWDfVciE7KCBCEpVbpJB9fbiTQ    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>        emSdpR8owrjpCt0byLf6OSunFEI8wDSC2jgAaSTLVq8ojNZAJskNCdk3bk-XjZhFo0xxgODGrbSjo3CjvJN3jmnSqt1Leq6r4SM-eXxLwHAv9Nmkj-E45JA2sd6MGovmmaQ1if6K5CIIqEXi6pZoohLLslyDTsdTBOb1AQ8eKArMyOhdz6dxBX8RKSATw5FhnVyeshJAhpB5sFopPk97aPkJffl2PXC7CRo14E85icaFUvWQT1MrRNs0XMRBlwatthSYNQvIHq8ahzh3j0XnomgPOdXBZe3soYUaHJs_ZT0JVq8dADNVYIOHux5weB_1nNIC9iRDTbjXLFZdyhZF6Q    </li></ul>Make a request.<p><b>GET /v1/api/cards</b> request<br>    Important - when including query parameters into any of GET requests, they should be preceded by a \"?\" after the URI    in the JWS Payload and combined with \"&\", if multiple are present.<p>For example:</p><ul>    <li>        <b>uri</b> -<br> /v1/api/cards    </li>    <li>        <b>query parameters</b>:<br>        cardAccount: \"eq:00000000\",<br>        cardName: \"eq:Test Test\",<br>        count: \"1000\",<br>        product: \"eq:999\"    </li>    <li>        <b>Resulting URI and query parameters</b> (token-body):<br>        /v1/api/cards?cardAccount=eq:00000000&cardName=eq:Test Test&count=1000&product=eq:999    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>        308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100B73748751B973456931B380163CE168BF884B26B1D27FA6978ACCEF829DF2DBA5D2845ABE997E82531FFD1C22D5E9F1ED6672F6D75D6572B688AE24AB9ED5138D5B5AA484FA04CEC393DB84C833BDDF7507BA1F5DC59030BE6E0FAC16613FADF3D06123621C11FB8F64FE777B9D0308DCEDCBAF78EB76FDAA4151C96E48B68507FA254AEFE44F0C0D103CAFE2A4ADB676E2B4FE26DB4AD671B8F02E74FE4B7E36156F44AA7CCE39DB0879926CEE2EF43502C7F99A1A3DDE72ECAC39E83F7B6819F2C1B8606537D6A04A1D6CC1FF0044F00360EA2F899BD9D9C6D748B20E2EB17CF908D6117B6B5B97490F5D68DC2D039B11EB2D920B8CEE9EFF10D0408A69971020301000102820100584078D67003D5621E59EE103D52F7E9543C16F521863048BDA7FFC9E9E166D56E2A07E5570FA4F2C5B2C40714738F91FB1498F3D0DEFDFE1ACD4B53535BCCD3E39D2DF8C6E2202B692721AF39478D13A3E0E992D420CC26DF7F5F49E9319696117EEB26C7FB8E9C398923A5B80B6057EE5CC4729A7C2194DF948BC0E3358378F1E3381BB62FE98C9372BB6EBBC24E8670E80EF991173D6D6347ADF1F8C9B67920D88072978D025F16F50273EDA0F0AA6FABC6F9687259A21142FD1E5C3DE99EE949DF1E81C7E7CDA4DD806A5CE2105DDC0CF7A2274913971C234D04D7C577C76F2FAE75379A0582FFA0333A01A7BE29ECA203AD7E3E71271EDEFDC0069F684102818100E7C26736089B4FB37BE4873EB2A726BBCD9CA3A3F3EB6BAD1433127614ABC7E6763D118A4A8DB45BBE31A67A853A44BA21EDF50533ACD348A0A9C1F648EE510426ECC845C01185C6793E9D4866281B9C01AE8BA96D0897BCBB0F638696CD6935794506AA3EC14BAB1A7F68DDF70DA21C70249FAD3FE88AB4C6708691327E4BC302818100CA611446E3A9A1997D175BDEADDE02B5C09E3E76024741F2DAAEF8689CF5DCA1C5FA9AE8A6F2F09E837E9D181C5607374DFBAA53B02C3908DCFA1817DEF1797F64743218B1616B55125E618B45F11003BA94D89E357358BD54A1DA324D893FDD9C0B30428A770F410B061C011CDB541E344B4C01B17CB615A1DCAD1F5EEB96BB02818100E40750C9C75A18E73E053254AC2EEE5B6608B2B18433A4341D65CACA47B864ED0A7537A6DB87E56747114EFDC9CBF507368F0CBF5B82B638056C419D73509881FF528612AAD212CF9F47CE3507DE7A9BDAC3C442A5370924F6E0434A8F61F81C56FF65796859837C0C8C43BFF16E868C78827061643A070FAB17D82F508117450281804FE4A32C99138E4819A9EF0AA978CB7914E163A7129F2ED9C09AF255DA20F548A    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>        WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br> {\"x5t#S256\":\"WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b> -<br>        /v1/api/cardscardAccount=eq:00000000&cardName=eq:Test Test&count=1000&product=eq:999    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>        eyJhbGciOiJSUzI1NiIsIng1dCNTMjU2IjoiV09RRUV5cmNIYVZvYkhNUE9SbXVNZTlpQUM1Z1NJMFRwWHlJZkExbEtJMCJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>        L3YxL2FwaS9jYXJkc2NhcmRBY2NvdW50PWVxOjAwMDAwMDAwJmNhcmROYW1lPWVxOlRlc3QgVGVzdCZjb3VudD0xMDAwJnByb2R1Y3Q9ZXE6OTk5.A2wnCqLNRQi81pzpVvTKDXI4GTMWZ-fjxEt7usuuVVh7BGcEEvRZ4Giy-mLU1e4jBsH5QR6fGT4jLOuTQ_4ItDChH5IBJNHDGTDBoXNNWYFrZS5bgFdtt7--U9i-jPsGySue_iT87Vy8ReJ147mberWN8S18u630gCamhdMdAwPmxIaD_4l1YSKHt29Z_5bK_-uvMIiLrV2ncqRZ_SCnna-GwCceTSa8IKn4uHqb6qngIJYxucWHhye-h-f26KGdhzoTbwz-2u90zNZ0pB3NESclQnmfdKbj8CzY-8Chr9GrAx4p4JLXh12rVGWUTojfsW2rzVgV_H3aw-bCDKVt1g    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>        A2wnCqLNRQi81pzpVvTKDXI4GTMWZ-fjxEt7usuuVVh7BGcEEvRZ4Giy-mLU1e4jBsH5QR6fGT4jLOuTQ_4ItDChH5IBJNHDGTDBoXNNWYFrZS5bgFdtt7--U9i-jPsGySue_iT87Vy8ReJ147mberWN8S18u630gCamhdMdAwPmxIaD_4l1YSKHt29Z_5bK_-uvMIiLrV2ncqRZ_SCnna-GwCceTSa8IKn4uHqb6qngIJYxucWHhye-h-f26KGdhzoTbwz-2u90zNZ0pB3NESclQnmfdKbj8CzY-8Chr9GrAx4p4JLXh12rVGWUTojfsW2rzVgV_H3aw-bCDKVt1g    </li></ul>Make a request.<p>    <b>If you have further questions regarding DECTA API, please contact our DECTA Support Team -</b><br></p>

API version: 2.11
Contact: support@decta.com
*/

package decta_api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// CardsAPIService CardsAPI service
type CardsAPIService service

type ApiAssignPinRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	tokenHeader    *string
	tokenSignature *string
	pinAssign      *PinAssign
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiAssignPinRequest) TokenHeader(tokenHeader string) ApiAssignPinRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiAssignPinRequest) TokenSignature(tokenSignature string) ApiAssignPinRequest {
	r.tokenSignature = &tokenSignature
	return r
}

func (r ApiAssignPinRequest) PinAssign(pinAssign PinAssign) ApiAssignPinRequest {
	r.pinAssign = &pinAssign
	return r
}

// Request ID (UUID format)
func (r ApiAssignPinRequest) RequestId(requestId string) ApiAssignPinRequest {
	r.requestId = &requestId
	return r
}

func (r ApiAssignPinRequest) Execute() (*http.Response, error) {
	return r.ApiService.AssignPinExecute(r)
}

/*
AssignPin Change PIN code

The request allows cardholder to change PIN code.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiAssignPinRequest
*/
func (a *CardsAPIService) AssignPin(ctx context.Context, ppan string) ApiAssignPinRequest {
	return ApiAssignPinRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
func (a *CardsAPIService) AssignPinExecute(r ApiAssignPinRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.AssignPin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/assign-pin"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return nil, reportError("tokenSignature is required and must be specified")
	}
	if r.pinAssign == nil {
		return nil, reportError("pinAssign is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	// body params
	localVarPostBody = r.pinAssign
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCardData1Request struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	param1         *string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// parameter 1
func (r ApiGetCardData1Request) Param1(param1 string) ApiGetCardData1Request {
	r.param1 = &param1
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetCardData1Request) TokenHeader(tokenHeader string) ApiGetCardData1Request {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetCardData1Request) TokenSignature(tokenSignature string) ApiGetCardData1Request {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetCardData1Request) RequestId(requestId string) ApiGetCardData1Request {
	r.requestId = &requestId
	return r
}

func (r ApiGetCardData1Request) Execute() (*CardData1ApiDto, *http.Response, error) {
	return r.ApiService.GetCardData1Execute(r)
}

/*
GetCardData1 Get sensitive card information 1

The request returns users sensitive card data like PIN code in encrypted format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetCardData1Request
*/
func (a *CardsAPIService) GetCardData1(ctx context.Context, ppan string) ApiGetCardData1Request {
	return ApiGetCardData1Request{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardData1ApiDto
func (a *CardsAPIService) GetCardData1Execute(r ApiGetCardData1Request) (*CardData1ApiDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardData1ApiDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetCardData1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/card-data1"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.param1 == nil {
		return localVarReturnValue, nil, reportError("param1 is required and must be specified")
	}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "param1", r.param1, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCardData2Request struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	param1         *string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// parameter 1
func (r ApiGetCardData2Request) Param1(param1 string) ApiGetCardData2Request {
	r.param1 = &param1
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetCardData2Request) TokenHeader(tokenHeader string) ApiGetCardData2Request {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetCardData2Request) TokenSignature(tokenSignature string) ApiGetCardData2Request {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetCardData2Request) RequestId(requestId string) ApiGetCardData2Request {
	r.requestId = &requestId
	return r
}

func (r ApiGetCardData2Request) Execute() (*CardData2ApiDto, *http.Response, error) {
	return r.ApiService.GetCardData2Execute(r)
}

/*
GetCardData2 Get sensitive card information 2

The request returns users with sensitive card data like full PAN and CVV in encrypted format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetCardData2Request
*/
func (a *CardsAPIService) GetCardData2(ctx context.Context, ppan string) ApiGetCardData2Request {
	return ApiGetCardData2Request{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardData2ApiDto
func (a *CardsAPIService) GetCardData2Execute(r ApiGetCardData2Request) (*CardData2ApiDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardData2ApiDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetCardData2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/card-data2"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.param1 == nil {
		return localVarReturnValue, nil, reportError("param1 is required and must be specified")
	}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "param1", r.param1, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCardData3Request struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	secret         *string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info
func (r ApiGetCardData3Request) Secret(secret string) ApiGetCardData3Request {
	r.secret = &secret
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetCardData3Request) TokenHeader(tokenHeader string) ApiGetCardData3Request {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetCardData3Request) TokenSignature(tokenSignature string) ApiGetCardData3Request {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetCardData3Request) RequestId(requestId string) ApiGetCardData3Request {
	r.requestId = &requestId
	return r
}

func (r ApiGetCardData3Request) Execute() (*CardData3HolderApiDto, *http.Response, error) {
	return r.ApiService.GetCardData3Execute(r)
}

/*
GetCardData3 Get sensitive card information for tokenization-3

The request allows users to receive card data including clear pan, cardholder name, and expiry date for the token service provider.
To ensure an additional layer of security it is mandatory to use the OTP Secret.
The OTP secret allows the card issuer to delegate permissions for a token service provider to fetch the sensitive data for a particular PPAN.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetCardData3Request
*/
func (a *CardsAPIService) GetCardData3(ctx context.Context, ppan string) ApiGetCardData3Request {
	return ApiGetCardData3Request{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardData3HolderApiDto
func (a *CardsAPIService) GetCardData3Execute(r ApiGetCardData3Request) (*CardData3HolderApiDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardData3HolderApiDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetCardData3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/card-data3"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.secret == nil {
		return localVarReturnValue, nil, reportError("secret is required and must be specified")
	}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "secret", r.secret, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCardData4Request struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	secret         *string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// OTP secret value generated by card issuer and passed to the token service provider to grant permission to fetch for the sensitive info
func (r ApiGetCardData4Request) Secret(secret string) ApiGetCardData4Request {
	r.secret = &secret
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetCardData4Request) TokenHeader(tokenHeader string) ApiGetCardData4Request {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetCardData4Request) TokenSignature(tokenSignature string) ApiGetCardData4Request {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetCardData4Request) RequestId(requestId string) ApiGetCardData4Request {
	r.requestId = &requestId
	return r
}

func (r ApiGetCardData4Request) Execute() (*Data4, *http.Response, error) {
	return r.ApiService.GetCardData4Execute(r)
}

/*
GetCardData4 Get sensitive card information for tokenization-4

The request allows users to receive card data including clear pan, cardholder name, CVC, and expiry date for the third-party data service provider. As an additional layer of security, it is necessary to use the OTP Secret. OTP secret allows card issuers to delegate permissions for third-party service providers to fetch sensitive data for a particular PPAN.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetCardData4Request
*/
func (a *CardsAPIService) GetCardData4(ctx context.Context, ppan string) ApiGetCardData4Request {
	return ApiGetCardData4Request{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return Data4
func (a *CardsAPIService) GetCardData4Execute(r ApiGetCardData4Request) (*Data4, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Data4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetCardData4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/card-data4"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.secret == nil {
		return localVarReturnValue, nil, reportError("secret is required and must be specified")
	}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "secret", r.secret, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCardData5Request struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	param1         *string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// parameter 1
func (r ApiGetCardData5Request) Param1(param1 string) ApiGetCardData5Request {
	r.param1 = &param1
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetCardData5Request) TokenHeader(tokenHeader string) ApiGetCardData5Request {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetCardData5Request) TokenSignature(tokenSignature string) ApiGetCardData5Request {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetCardData5Request) RequestId(requestId string) ApiGetCardData5Request {
	r.requestId = &requestId
	return r
}

func (r ApiGetCardData5Request) Execute() (*CardData2ApiDto, *http.Response, error) {
	return r.ApiService.GetCardData5Execute(r)
}

/*
GetCardData5 Get sensitive card information 5

The request returns users with sensitive card data like full PAN and CVV in encrypted format .

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetCardData5Request
*/
func (a *CardsAPIService) GetCardData5(ctx context.Context, ppan string) ApiGetCardData5Request {
	return ApiGetCardData5Request{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardData2ApiDto
func (a *CardsAPIService) GetCardData5Execute(r ApiGetCardData5Request) (*CardData2ApiDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardData2ApiDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetCardData5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/card-data5"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.param1 == nil {
		return localVarReturnValue, nil, reportError("param1 is required and must be specified")
	}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "param1", r.param1, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInfoRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetInfoRequest) TokenHeader(tokenHeader string) ApiGetInfoRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetInfoRequest) TokenSignature(tokenSignature string) ApiGetInfoRequest {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetInfoRequest) RequestId(requestId string) ApiGetInfoRequest {
	r.requestId = &requestId
	return r
}

func (r ApiGetInfoRequest) Execute() (*CardInfo, *http.Response, error) {
	return r.ApiService.GetInfoExecute(r)
}

/*
GetInfo Get card information

The request allows users to get card information with available balance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetInfoRequest
*/
func (a *CardsAPIService) GetInfo(ctx context.Context, ppan string) ApiGetInfoRequest {
	return ApiGetInfoRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardInfo
func (a *CardsAPIService) GetInfoExecute(r ApiGetInfoRequest) (*CardInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetListRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	tokenHeader    *string
	tokenSignature *string
	cardName       *string
	expiry         *[]DateFilter
	cardState      *string
	cardAccount    *string
	clientId       *string
	product        *string
	count          *int32
	startId        *int64
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetListRequest) TokenHeader(tokenHeader string) ApiGetListRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetListRequest) TokenSignature(tokenSignature string) ApiGetListRequest {
	r.tokenSignature = &tokenSignature
	return r
}

// The filter allows users to get a list of cards by the selected name on the card. Allowed filter operator(s): eq, ne
func (r ApiGetListRequest) CardName(cardName string) ApiGetListRequest {
	r.cardName = &cardName
	return r
}

// The filter allows users to sort cards by the selected expiration date in the list. Allowed filter operator(s): eq, ne, gt, ge, lt, le
func (r ApiGetListRequest) Expiry(expiry []DateFilter) ApiGetListRequest {
	r.expiry = &expiry
	return r
}

// The filter allows users to get a list of cards by the selected  card status. Allowed filter operator(s): eq
func (r ApiGetListRequest) CardState(cardState string) ApiGetListRequest {
	r.cardState = &cardState
	return r
}

// The filter allows users to get a list of cards by the selected card account. Allowed filter operator(s): eq, ne
func (r ApiGetListRequest) CardAccount(cardAccount string) ApiGetListRequest {
	r.cardAccount = &cardAccount
	return r
}

// The filter allows users to get a list of cards by the selected client identification number. Allowed filter operator(s): eq, ne
func (r ApiGetListRequest) ClientId(clientId string) ApiGetListRequest {
	r.clientId = &clientId
	return r
}

// The filter allows users to get a list of cards by selected product code. Allowed filter operator(s): eq, ne
func (r ApiGetListRequest) Product(product string) ApiGetListRequest {
	r.product = &product
	return r
}

// Number of cards in the list
func (r ApiGetListRequest) Count(count int32) ApiGetListRequest {
	r.count = &count
	return r
}

// Start the ordered list from card with selected id
func (r ApiGetListRequest) StartId(startId int64) ApiGetListRequest {
	r.startId = &startId
	return r
}

// Request ID (UUID format)
func (r ApiGetListRequest) RequestId(requestId string) ApiGetListRequest {
	r.requestId = &requestId
	return r
}

func (r ApiGetListRequest) Execute() (*CardInfoDataArray, *http.Response, error) {
	return r.ApiService.GetListExecute(r)
}

/*
GetList Search cards

The request allows users to get a list of cards. This request returns partial information about cards, to receive detailed information about a certain card, please use a specific card GET card request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetListRequest
*/
func (a *CardsAPIService) GetList(ctx context.Context) ApiGetListRequest {
	return ApiGetListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CardInfoDataArray
func (a *CardsAPIService) GetListExecute(r ApiGetListRequest) (*CardInfoDataArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardInfoDataArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	if r.cardName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardName", r.cardName, "form", "")
	}
	if r.expiry != nil {
		t := *r.expiry
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "expiry", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "expiry", t, "form", "multi")
		}
	}
	if r.cardState != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardState", r.cardState, "form", "")
	}
	if r.cardAccount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cardAccount", r.cardAccount, "form", "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "form", "")
	}
	if r.product != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product", r.product, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 10
		r.count = &defaultValue
	}
	if r.startId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startId", r.startId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTspSecretRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetTspSecretRequest) TokenHeader(tokenHeader string) ApiGetTspSecretRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetTspSecretRequest) TokenSignature(tokenSignature string) ApiGetTspSecretRequest {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiGetTspSecretRequest) RequestId(requestId string) ApiGetTspSecretRequest {
	r.requestId = &requestId
	return r
}

func (r ApiGetTspSecretRequest) Execute() (*TspSecret, *http.Response, error) {
	return r.ApiService.GetTspSecretExecute(r)
}

/*
GetTspSecret Get OTP secret for tokenization

The request returns the OTP secret value to DAPI users. OTP secret should be passed to the token_service_provider role for authorization to the sensitive data in card data 3 and card data 4 requests.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiGetTspSecretRequest
*/
func (a *CardsAPIService) GetTspSecret(ctx context.Context, ppan string) ApiGetTspSecretRequest {
	return ApiGetTspSecretRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return TspSecret
func (a *CardsAPIService) GetTspSecretExecute(r ApiGetTspSecretRequest) (*TspSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TspSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.GetTspSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/otp-secret"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRenewCardRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiRenewCardRequest) TokenHeader(tokenHeader string) ApiRenewCardRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiRenewCardRequest) TokenSignature(tokenSignature string) ApiRenewCardRequest {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiRenewCardRequest) RequestId(requestId string) ApiRenewCardRequest {
	r.requestId = &requestId
	return r
}

func (r ApiRenewCardRequest) Execute() (*CardRenewInfo, *http.Response, error) {
	return r.ApiService.RenewCardExecute(r)
}

/*
RenewCard Renew card

The request allows users to renew cards by selected PPAN

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiRenewCardRequest
*/
func (a *CardsAPIService) RenewCard(ctx context.Context, ppan string) ApiRenewCardRequest {
	return ApiRenewCardRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardRenewInfo
func (a *CardsAPIService) RenewCardExecute(r ApiRenewCardRequest) (*CardRenewInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardRenewInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.RenewCard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/renew"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReplaceCardRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	tokenHeader    *string
	tokenSignature *string
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiReplaceCardRequest) TokenHeader(tokenHeader string) ApiReplaceCardRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiReplaceCardRequest) TokenSignature(tokenSignature string) ApiReplaceCardRequest {
	r.tokenSignature = &tokenSignature
	return r
}

// Request ID (UUID format)
func (r ApiReplaceCardRequest) RequestId(requestId string) ApiReplaceCardRequest {
	r.requestId = &requestId
	return r
}

func (r ApiReplaceCardRequest) Execute() (*CardInfo, *http.Response, error) {
	return r.ApiService.ReplaceCardExecute(r)
}

/*
ReplaceCard Replace card

The request allows users to replace cards by selected PPAN

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiReplaceCardRequest
*/
func (a *CardsAPIService) ReplaceCard(ctx context.Context, ppan string) ApiReplaceCardRequest {
	return ApiReplaceCardRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
//
//	@return CardInfo
func (a *CardsAPIService) ReplaceCardExecute(r ApiReplaceCardRequest) (*CardInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CardInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.ReplaceCard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/replace"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCardUserDefinedFieldsRequest struct {
	ctx                   context.Context
	ApiService            *CardsAPIService
	ppan                  string
	tokenHeader           *string
	tokenSignature        *string
	cardUserDefinedFields *CardUserDefinedFields
	requestId             *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiUpdateCardUserDefinedFieldsRequest) TokenHeader(tokenHeader string) ApiUpdateCardUserDefinedFieldsRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiUpdateCardUserDefinedFieldsRequest) TokenSignature(tokenSignature string) ApiUpdateCardUserDefinedFieldsRequest {
	r.tokenSignature = &tokenSignature
	return r
}

func (r ApiUpdateCardUserDefinedFieldsRequest) CardUserDefinedFields(cardUserDefinedFields CardUserDefinedFields) ApiUpdateCardUserDefinedFieldsRequest {
	r.cardUserDefinedFields = &cardUserDefinedFields
	return r
}

// Request ID (UUID format)
func (r ApiUpdateCardUserDefinedFieldsRequest) RequestId(requestId string) ApiUpdateCardUserDefinedFieldsRequest {
	r.requestId = &requestId
	return r
}

func (r ApiUpdateCardUserDefinedFieldsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateCardUserDefinedFieldsExecute(r)
}

/*
UpdateCardUserDefinedFields Change card user defined fields

The request allows users to change card user defined fields

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan Masked card number
	@return ApiUpdateCardUserDefinedFieldsRequest
*/
func (a *CardsAPIService) UpdateCardUserDefinedFields(ctx context.Context, ppan string) ApiUpdateCardUserDefinedFieldsRequest {
	return ApiUpdateCardUserDefinedFieldsRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
func (a *CardsAPIService) UpdateCardUserDefinedFieldsExecute(r ApiUpdateCardUserDefinedFieldsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.UpdateCardUserDefinedFields")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/user-defined-fields"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.ppan) < 16 {
		return nil, reportError("ppan must have at least 16 elements")
	}
	if strlen(r.ppan) > 16 {
		return nil, reportError("ppan must have less than 16 elements")
	}
	if r.tokenHeader == nil {
		return nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return nil, reportError("tokenSignature is required and must be specified")
	}
	if r.cardUserDefinedFields == nil {
		return nil, reportError("cardUserDefinedFields is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	// body params
	localVarPostBody = r.cardUserDefinedFields
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateStateRequest struct {
	ctx            context.Context
	ApiService     *CardsAPIService
	ppan           string
	tokenHeader    *string
	tokenSignature *string
	statusChange   *StatusChange
	requestId      *string
}

// URL64Encoded without padding Header part of JWS token
func (r ApiUpdateStateRequest) TokenHeader(tokenHeader string) ApiUpdateStateRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiUpdateStateRequest) TokenSignature(tokenSignature string) ApiUpdateStateRequest {
	r.tokenSignature = &tokenSignature
	return r
}

func (r ApiUpdateStateRequest) StatusChange(statusChange StatusChange) ApiUpdateStateRequest {
	r.statusChange = &statusChange
	return r
}

// Request ID (UUID format)
func (r ApiUpdateStateRequest) RequestId(requestId string) ApiUpdateStateRequest {
	r.requestId = &requestId
	return r
}

func (r ApiUpdateStateRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateStateExecute(r)
}

/*
UpdateState Change card status

The request allows users to block, unblock and activate cards for selected PPAN. Once the card is BLOCKED_PERMANENTLY it can't be unblocked.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ppan masked card number
	@return ApiUpdateStateRequest
*/
func (a *CardsAPIService) UpdateState(ctx context.Context, ppan string) ApiUpdateStateRequest {
	return ApiUpdateStateRequest{
		ApiService: a,
		ctx:        ctx,
		ppan:       ppan,
	}
}

// Execute executes the request
func (a *CardsAPIService) UpdateStateExecute(r ApiUpdateStateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsAPIService.UpdateState")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api/cards/{ppan}/state"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenHeader == nil {
		return nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return nil, reportError("tokenSignature is required and must be specified")
	}
	if r.statusChange == nil {
		return nil, reportError("statusChange is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	if r.requestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	}
	// body params
	localVarPostBody = r.statusChange
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
