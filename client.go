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
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

// APIClient manages communication with the Decta API API v2.11
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccountsAPI *AccountsAPIService

	AuthorizationsAPI *AuthorizationsAPIService

	CardOrderAPI *CardOrderAPIService

	CardsAPI *CardsAPIService

	ClientsAPI *ClientsAPIService

	DectaSecureAPI *DectaSecureAPIService

	LimitsAPI *LimitsAPIService

	RemindersAPI *RemindersAPIService

	TokenAPI *TokenAPIService

	TransactionsAPI *TransactionsAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccountsAPI = (*AccountsAPIService)(&c.common)
	c.AuthorizationsAPI = (*AuthorizationsAPIService)(&c.common)
	c.CardOrderAPI = (*CardOrderAPIService)(&c.common)
	c.CardsAPI = (*CardsAPIService)(&c.common)
	c.ClientsAPI = (*ClientsAPIService)(&c.common)
	c.DectaSecureAPI = (*DectaSecureAPIService)(&c.common)
	c.LimitsAPI = (*LimitsAPIService)(&c.common)
	c.RemindersAPI = (*RemindersAPIService)(&c.common)
	c.TokenAPI = (*TokenAPIService)(&c.common)
	c.TransactionsAPI = (*TransactionsAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString(obj interface{}, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param, ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
				return
			}
			if t, ok := obj.(time.Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
				return
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			var lenIndValue = indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				var arrayValue = indValue.Index(i)
				var keyPrefixForCollectionType = keyPrefix
				if style == "deepObject" {
					keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
			}
			return

		case reflect.Map:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
		break
	case map[string]string:
		valuesMap[keyPrefix] = value
		break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
	fileBytes    []byte
	fileName     string
	formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
					return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
					return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	localVarRequest, err = c.cfg.PrepareRequest(localVarRequest)
	if err != nil {
		return nil, err
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
