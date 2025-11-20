# Go API client for dectav2

<h2>Introduction</h2><p>This is a documented guide that    describes the process of DECTA card and token operations for both private and business clients. As a new or already    experienced API user you will find in this documentation all necessary and required information to start using DECTA    API and improve your experience with DECTA.</p><h2>Authentication and Authorisation</h2><p>Each request to DECTA API must be signed with a certificate that allows DECTA to identify API clients. This    guarantees the authenticity of the request received from the client.</p><h3>Authentication</h3><p>The authentication process is executed only once for each client.<br>    On successful authentication DECTA assigns a certificate to the client which must be used in every API    request in the future.</p><p>    To access a detailed description of the authentication process, choose <b>Service</b> specification at the top right    corner or click <a        href=\"/?urls.primaryName=SERVICE\" target=\"_blank\">here</a>.</p><h3>Authorisation</h3><p>    Each request must be signed using Jose specifications in particular - <a        href=\"https://tools.ietf.org/html/rfc7515\" target=\"_blank\">RFC7515</a>. It is vital for users to follow the    provided standard and use <b>Base64url    Encoding without Padding</b> for token encoding.</p><h3>Request with filters</h3><p>    Several requests support data filters. Filter syntax can be found below -<p>    <code>        &lt;requestFilter&gt; ::= &lt;filterName&gt;=[eq|ne|gt|ge|lt|le]:&lt;filterValue&gt;    </code></p><ul>Where:    <li>eq - equals</li>    <li>ne - not equals</li>    <li>gt - greater than</li>    <li>ge - greater or equals than</li>    <li>lt - less than</li>    <li>le - less or equals than</li></ul><ul>Example:    <li>cardState: \"eq:BLOCKED_BY_HOLDER\"</li></ul><h2>Changelog</h2>Information about the latest DECTA API changes and improvements can be found <a href=\"/changelog.html\" target=\"_blank\">here.</a>.<h4>Request examples</h4><p>    We recommend users follow the instructions in the next sections below and create the application step by step.</p><p>    To understand the process of DECTA API requests better, we have prepared two basic API calls as examples.</p><p>    <b>POST /v1/api/cards/order</b> request</p><p>For example:</p><ul>    <li>        <b>uri</b> - <br>/v1/api/cards/order    </li>    <li>        <b>body</b> - <br>{\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100C5BD899E1BD0E7D0C7F814976C449288738ECBD396D3BB7ABE8ED6ADACB60AA4622594FB2CE5C97F01274F3F350935A170013743B557882674E2C501C0A17943300C6E39D6D171B06980E12B8946050D4255F6E8E293FCF3DA848DE2486296EB24CA68619ED954804FF00F4936B4289AE04FCA3AD8B0AFA1D3BF7D0DD956D59B528BA284595116B512B5234A53B8CD746268F6F93503925B565DEF8BE9A9A69E0ECE869B40D11F1CF65FC92FDE8159943437F161ED33B4BB83FE5E8F3898E2C03847E84F92BB25CF7835A67F8CA6C24A12F81CC99CE7FBAC387F6D1232668DCA64B1CA0DBFF1292DB7A0C4D129267F6779826AE86337CE00EA789A3F4BCC0E8F0203010001028201001B42BD15320A83DAF1A68ED82D51EACDAAF8B5BAEEE133C8813AD9F7C986B342E451000FFCA99534D028E630F593D673F22808C169DFAE6276DCC8822C57568B5FB56700CC4FC597A783E0070B7EBEB58C3EA3B28982136E5EE2C558BBD7F3AC693F0A52E2FE5D3AEF4A4E212ECC7988988207844359606D33EF0343848207228B91AE9F37D02FD1F5D9994D0DBB9B15DE1F3BDE62980BC3645226896D2AC9A79DE49F0129E2F012D45BEBABA6AEFC7F0E09FB4C551C976BE2AD393EB1D7A448A6579547EA481F87B14EEDC0DB3FB1CAB326F5B4C7B7655BE8D96EE167BA64CFDA033A15F503BED8B62067A162C1ABEEF2B0719360531EB98E127BC21CCF7C1902818100EA311697B86CC3257E2C1BC9CE1928A1ECA2D9F4DC07A5ADA087D0E053AA1F9D3EB46C42627CE5541F386B8ED356033B08D25F195DD9AC27B944FEA16E16C7F8D68FFF43BC76C7B919CA0B6CEAC775E0D59FA55E108182406AFDC3A00C38BA7968A775604813679B0A1BA4389F21B1DDB8002A09FDF94E3875B4D971179645E502818100D8277B9B4BFDB22A2E671BE25BF41D07EBC13D7872C18BBC12C925914594580AB6E0B35BB479F7D437F35F1D5602A4EB04E3B4BA9F3FD66983AFE28D56B9945A55716CE655B6A51926714C97CDA211272A6D8D1A6A5DA903955DDAE4CCF5CB17AB579872ABFB3A4AD2A9B6B7781C217996753D8F08B7FEBACE0437E7E3127B630281804110A25644C6F2F0D7BD297B47E77582875C3771F02FAFB82D818E66C4D7DF30B0DB2FB8C1E43152CDD9BF084F6EF636D09CA20F4A23CFD0B98891B608000C4EF64DCB7DE63E99FEB34B8F211399970680537A7E47D7B37872414EB888ACBA224111456B1A9B498DA9A1445EF6D745B3740B97A26F0D74CDFAE5E403A7A9F88102818100A3196FB867A79F471ACA6C8CE9CB921846F901D41FEB10F031D09B4B39904CBA90F18E04B5728EE10FF2D0A3472229A4B884C4FB486C97F1F80EC048DBD69E2D676975BD7B919C9414A0CB25846134C633CA83FD19A36083B6B2221820DB9A8A5A9C770897E2B2428120D4212E3D33B9162192F899A094520A5FBE7B77A89B7702818100D9DD12BF56479E6875BF4B82450C926F07E56FCC4F8A54594E7C6D76FFBDE50DC1C57097665580D88A9E13989F230EBA30254C056DC664990E06FA54565A977A5693B847F1C3025FF733F24D1F135B51A9FEF2D08476AF26A129F230DFE04140995D0C5BD02FE62067B652822BD25AEDAAAFACE932981FE0FF38A94CE50C0B25    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br>{\"x5t#S256\":\"g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b>-<br>/v1/api/cards/order{\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>eyJ4NXQjUzI1NiI6Imc4R2tjNE9xWndJRFY3SFNla2xwUWNKWlpXeVRWa08xSktvZlh0QU42M0kiLCJhbGciOiJSUzI1NiJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>L3YxL2FwaS9jYXJkcy9vcmRlcnsiY2FyZCI6eyJhY2NvdW50T3duZXJSZWxhdGlvbiI6Ik9XTkVSIiwiY3VycmVuY2llcyI6WyJFVVIiXSwiZGVsaXZlcnlBZGRyZXNzIjp7ImNpdHkiOiJNYWRvbmEiLCJjb3VudHJ5IjoiTFZBIiwibmFtZSI6IlRlc3QiLCJwaG9uZSI6OTExMywic2hpcG1lbnQiOiJTVEFOREFSRCIsInN0cmVldCI6IjYyIE1hZG9uYSBzdHJlZXQiLCJzdXJuYW1lIjoiVGVzdFMiLCJ6aXBDb2RlIjoiTFYtMTAwMCJ9LCJob2xkZXIiOnsiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2EiLCJjb3VudHJ5IjoiTFZBIiwic3RyZWV0IjoiNyBTdHJlZXQiLCJ6aXBDb2RlIjoiTFYwMDAwIn0sImRvY3VtZW50Ijp7ImJpcnRoRGF0ZSI6IjE5NTAtMTItMzEiLCJjb3VudHJ5Q29kZSI6IkVVUiIsImV4cGlyeURhdGUiOiIyMDU2LTEyLTMxIiwiaXNzdWluZ0RhdGUiOiIyMDE2LTEyLTMxIiwibnVtYmVyIjoiVFRIb2xkZXIiLCJzdWJ0eXBlIjoiUkVTSURFTlRfUEVSTUlUX0lEIiwidHlwZSI6IkRSSVZJTkdfTElDRU5TRSJ9LCJlbWFpbCI6Im9ob2xkZXJAZGVjdGEuY29tIiwibGFuZ3VhZ2UiOiJFTiIsIm1haWRlbk5hbWUiOiJUZXN0TSIsIm1vYmlsZVBob25lIjoiMTIzNDU2NzgiLCJuYW1lIjoiVGVzdFMiLCJyb2xlIjoiT1dORVIiLCJzdXJuYW1lIjoiVGVzdFMiLCJ1c2VyRGVmaW5lZEZpZWxkMSI6IjIifSwicGFzc3BocmFzZSI6InRlc3QycGFzcyIsInByaW9yaXR5IjoiVVJHRU5UIiwicHJvZHVjdENvZGUiOiI3MDEiLCJzdXBwbGVtZW50YXJ5IjoiZmFsc2UifSwiZXh0ZXJuYWxJZCI6IlByaXZhdGVfTmV3IiwicHJpdmF0ZUNsaWVudCI6eyJjb21tZW50IjoiQWdyZWVtZW50IG51bS4gMTIzMTQzIiwiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2FDdXJyZW50IiwiY291bnRyeSI6IkxWQSIsIm1vbnRocyI6NCwic3RyZWV0IjoiNiBzdHJlZXQgc3RyZWV0IiwiemlwQ29kZSI6IldTY3VyMTAwOSJ9LCJkb2N1bWVudCI6eyJiaXJ0aERhdGUiOiIxOTUwLTEyLTMxIiwiY291bnRyeUNvZGUiOiJMVkEiLCJleHBpcnlEYXRlIjoiMjA5OS0xMi0zMSIsImlzc3VpbmdEYXRlIjoiMTk3MS0xMi0zMSIsIm51bWJlciI6IjEyMzQiLCJzdWJ0eXBlIjoiTkFUSU9OQUxfSUQiLCJ0eXBlIjoiSURfQ0FSRCJ9LCJlbWFpbCI6ImVtYWlsMTIzNCIsImxhbmd1YWdlIjoiZW4iLCJtYWlkZW5OYW1lIjoiVGVzdCIsIm1vYmlsZVBob25lIjoiMzcxMjg4NjY2OTQiLCJuYW1lIjoiVGVzdCIsInBlcCI6Ik5PIiwic3VybmFtZSI6IlByaXZhdGVCYXNlTmV3IiwidXNlckRlZmluZWRGaWVsZDEiOiIyMDAwNjY2In19    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>wcVPkihLADBs774wPOyqUoK05tj6tSeFWaKuH1jRNED5pABPDpvsg7aDOtbBmORJGF0HfyVfquui5wRKuMQCUW2OaQe26Tlh194lxdfrcg_e1S3P-ZqqI6LEHKLGZFGSuuzT95Da5osSmfMTkr_LOQzT_vPiASr0KFSNSZtL60l2q3dUeZ2AcwTeVFi1Aga_K_mB5UeZUXPQf9XfQ1Eb2aU7Fl9YIsYL2M2B5NmQTgKfy5nuSQXhw3y_ddoiNFR1wHY-g6eX3oU1_kx6s5_GV7Jr9zftj6oCZjhOr-_Ohr2_UMxmpn1D7MGThhwlDguK7CfFWb_8nOvEqntYlTErDQ    </li></ul>Make a request.<p><b>GET /v1/api/cards</b> request<br>    Important - when including query parameters into any of GET requests, they should be preceded by a \"?\" after the URI    in the JWS Payload and combined with \"&\", if multiple are present.<p>For example:</p><ul>    <li>        <b>uri</b> -<br> /v1/api/cards    </li>    <li>        <b>query parameters</b>:<br>        cardAccount: \"eq:00000000\",<br>        cardName: \"eq:Test Test\",<br>        count: \"1000\",<br>        product: \"eq:999\"    </li>    <li>        <b>Resulting URI and query parameters</b>:<br>/v1/api/orders?product=eq:999&cardName=eq:Test        Test&count=1000&cardAccount=eq:00000000    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100C5BD899E1BD0E7D0C7F814976C449288738ECBD396D3BB7ABE8ED6ADACB60AA4622594FB2CE5C97F01274F3F350935A170013743B557882674E2C501C0A17943300C6E39D6D171B06980E12B8946050D4255F6E8E293FCF3DA848DE2486296EB24CA68619ED954804FF00F4936B4289AE04FCA3AD8B0AFA1D3BF7D0DD956D59B528BA284595116B512B5234A53B8CD746268F6F93503925B565DEF8BE9A9A69E0ECE869B40D11F1CF65FC92FDE8159943437F161ED33B4BB83FE5E8F3898E2C03847E84F92BB25CF7835A67F8CA6C24A12F81CC99CE7FBAC387F6D1232668DCA64B1CA0DBFF1292DB7A0C4D129267F6779826AE86337CE00EA789A3F4BCC0E8F0203010001028201001B42BD15320A83DAF1A68ED82D51EACDAAF8B5BAEEE133C8813AD9F7C986B342E451000FFCA99534D028E630F593D673F22808C169DFAE6276DCC8822C57568B5FB56700CC4FC597A783E0070B7EBEB58C3EA3B28982136E5EE2C558BBD7F3AC693F0A52E2FE5D3AEF4A4E212ECC7988988207844359606D33EF0343848207228B91AE9F37D02FD1F5D9994D0DBB9B15DE1F3BDE62980BC3645226896D2AC9A79DE49F0129E2F012D45BEBABA6AEFC7F0E09FB4C551C976BE2AD393EB1D7A448A6579547EA481F87B14EEDC0DB3FB1CAB326F5B4C7B7655BE8D96EE167BA64CFDA033A15F503BED8B62067A162C1ABEEF2B0719360531EB98E127BC21CCF7C1902818100EA311697B86CC3257E2C1BC9CE1928A1ECA2D9F4DC07A5ADA087D0E053AA1F9D3EB46C42627CE5541F386B8ED356033B08D25F195DD9AC27B944FEA16E16C7F8D68FFF43BC76C7B919CA0B6CEAC775E0D59FA55E108182406AFDC3A00C38BA7968A775604813679B0A1BA4389F21B1DDB8002A09FDF94E3875B4D971179645E502818100D8277B9B4BFDB22A2E671BE25BF41D07EBC13D7872C18BBC12C925914594580AB6E0B35BB479F7D437F35F1D5602A4EB04E3B4BA9F3FD66983AFE28D56B9945A55716CE655B6A51926714C97CDA211272A6D8D1A6A5DA903955DDAE4CCF5CB17AB579872ABFB3A4AD2A9B6B7781C217996753D8F08B7FEBACE0437E7E3127B630281804110A25644C6F2F0D7BD297B47E77582875C3771F02FAFB82D818E66C4D7DF30B0DB2FB8C1E43152CDD9BF084F6EF636D09CA20F4A23CFD0B98891B608000C4EF64DCB7DE63E99FEB34B8F211399970680537A7E47D7B37872414EB888ACBA224111456B1A9B498DA9A1445EF6D745B3740B97A26F0D74CDFAE5E403A7A9F88102818100A3196FB867A79F471ACA6C8CE9CB921846F901D41FEB10F031D09B4B39904CBA90F18E04B5728EE10FF2D0A3472229A4B884C4FB486C97F1F80EC048DBD69E2D676975BD7B919C9414A0CB25846134C633CA83FD19A36083B6B2221820DB9A8A5A9C770897E2B2428120D4212E3D33B9162192F899A094520A5FBE7B77A89B7702818100D9DD12BF56479E6875BF4B82450C926F07E56FCC4F8A54594E7C6D76FFBDE50DC1C57097665580D88A9E13989F230EBA30254C056DC664990E06FA54565A977A5693B847F1C3025FF733F24D1F135B51A9FEF2D08476AF26A129F230DFE04140995D0C5BD02FE62067B652822BD25AEDAAAFACE932981FE0FF38A94CE50C0B25    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br>{\"x5t#S256\":\"g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b> -<br>/v1/api/ordersproduct=eq:999&cardName=eq:Test Test&count=1000&cardAccount=eq:00000000    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>eyJ4NXQjUzI1NiI6Imc4R2tjNE9xWndJRFY3SFNla2xwUWNKWlpXeVRWa08xSktvZlh0QU42M0kiLCJhbGciOiJSUzI1NiJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>L3YxL2FwaS9vcmRlcnNwcm9kdWN0PWVxOjk5OSZjYXJkTmFtZT1lcTpUZXN0IFRlc3QmY291bnQ9MTAwMCZjYXJkQWNjb3VudD1lcTowMDAwMDAwMA    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>Iw8oJuOZvzxDmCjNZgOSHC-dr_y6hUUTHw8nT8C_bv6aYS9COLOLoTg3ANqo16RrBJM26hrprwMvVjEy1sz2-HiBTnZS6-HQo-KlqvCPqMH7HwSSbwFvedzcjAm2SC1xFUkGnAl5SJ_hxRWAejefmK4X3_ToJx3AN9JbJU40IeRCShKKSrNgui0A_zzkvza-JNiRhzzxRofrNxuUhF_xfGrOiIsGpG3k0l24bjzr-5YRf9vwGhQZRA93_ZAUUoLZ4pKi6CYcBd0dnmTLuIiFIpvBXKt9a-EZb5QJFk8GJnZ7pTXKRuJev2nW76rFhlC6rQB_RVoQTzsaIHae7swycg    </li></ul>Make a request.<p>    <b>If you have further questions regarding DECTA API, please contact our DECTA Support Team -</b><br></p>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.16
- Package version: 1.0.0
- Generator version: 7.14.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://decta.atlassian.net/servicedesk/customer/portal/2](https://decta.atlassian.net/servicedesk/customer/portal/2)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import dectav2 "github.com/GIT_USER_ID/GIT_REPO_ID/dectav2"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `dectav2.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), dectav2.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `dectav2.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), dectav2.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `dectav2.ContextOperationServerIndices` and `dectav2.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), dectav2.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dectav2.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost:8443*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Class10ClickToPayAPI* | [**AddClickToPay**](docs/Class10ClickToPayAPI.md#addclicktopay) | **Post** /v2/api/click-to-pay/{ppan} | Add card to Click to Pay
*Class10ClickToPayAPI* | [**ChangeClickToPay**](docs/Class10ClickToPayAPI.md#changeclicktopay) | **Patch** /v2/api/click-to-pay/{ppan} | Change Click to Pay information using card
*Class10ClickToPayAPI* | [**DeleteClickToPay**](docs/Class10ClickToPayAPI.md#deleteclicktopay) | **Delete** /v2/api/click-to-pay/{ppan} | Delete data from Click to Pay using card
*Class10ClickToPayAPI* | [**GetClickToPay**](docs/Class10ClickToPayAPI.md#getclicktopay) | **Get** /v2/api/click-to-pay/{clientId} | Get Click to Pay information using client id
*Class1CardOrderAPI* | [**GetOrdersList**](docs/Class1CardOrderAPI.md#getorderslist) | **Get** /v1/api/orders | Order status
*Class1CardOrderAPI* | [**OrderCard**](docs/Class1CardOrderAPI.md#ordercard) | **Post** /v1/api/cards/order | Order card
*Class1CardOrderAPI* | [**OrderGiftCard**](docs/Class1CardOrderAPI.md#ordergiftcard) | **Post** /v1/api/cards/order-gift-card | Create gift cards
*Class2CardsAPI* | [**AssignPin**](docs/Class2CardsAPI.md#assignpin) | **Post** /v1/api/cards/{ppan}/assign-pin | Change PIN code
*Class2CardsAPI* | [**Create**](docs/Class2CardsAPI.md#create) | **Post** /v2/api/cards | Create card V2
*Class2CardsAPI* | [**Create1**](docs/Class2CardsAPI.md#create1) | **Post** /v1/api/cards | Create card
*Class2CardsAPI* | [**CreateBatchCard**](docs/Class2CardsAPI.md#createbatchcard) | **Post** /v2/api/cards/batch | Batch create card
*Class2CardsAPI* | [**GetCardData1**](docs/Class2CardsAPI.md#getcarddata1) | **Get** /v1/api/cards/{ppan}/card-data1 | Get sensitive card information 1
*Class2CardsAPI* | [**GetCardData2**](docs/Class2CardsAPI.md#getcarddata2) | **Get** /v1/api/cards/{ppan}/card-data2 | Get sensitive card information 2
*Class2CardsAPI* | [**GetCardData3**](docs/Class2CardsAPI.md#getcarddata3) | **Get** /v1/api/cards/{ppan}/card-data3 | Get sensitive card information 3
*Class2CardsAPI* | [**GetCardData4**](docs/Class2CardsAPI.md#getcarddata4) | **Get** /v1/api/cards/{ppan}/card-data4 | Get sensitive card information 4
*Class2CardsAPI* | [**GetCardData5**](docs/Class2CardsAPI.md#getcarddata5) | **Get** /v1/api/cards/{ppan}/card-data5 | Get sensitive card information 5
*Class2CardsAPI* | [**GetCardData6**](docs/Class2CardsAPI.md#getcarddata6) | **Get** /v2/api/cards/{ppan}/card-data6 | Get sensitive card information 6
*Class2CardsAPI* | [**GetInfo**](docs/Class2CardsAPI.md#getinfo) | **Get** /v1/api/cards/{ppan} | Get card information
*Class2CardsAPI* | [**GetList**](docs/Class2CardsAPI.md#getlist) | **Get** /v1/api/cards | Search cards
*Class2CardsAPI* | [**GetTspSecret**](docs/Class2CardsAPI.md#gettspsecret) | **Get** /v1/api/cards/{ppan}/otp-secret | Get OTP secret
*Class2CardsAPI* | [**RenewCard**](docs/Class2CardsAPI.md#renewcard) | **Post** /v1/api/cards/{ppan}/renew | Renew card
*Class2CardsAPI* | [**ReplaceCard**](docs/Class2CardsAPI.md#replacecard) | **Post** /v1/api/cards/{ppan}/replace | Replace card
*Class2CardsAPI* | [**UpdateCardUserDefinedFields**](docs/Class2CardsAPI.md#updatecarduserdefinedfields) | **Patch** /v1/api/cards/{ppan}/user-defined-fields | Change card user defined fields
*Class2CardsAPI* | [**UpdateState**](docs/Class2CardsAPI.md#updatestate) | **Patch** /v1/api/cards/{ppan}/state | Change card status
*Class3RemindersAPI* | [**GetPinTryCounter**](docs/Class3RemindersAPI.md#getpintrycounter) | **Get** /v1/api/cards/{ppan}/pin-try-counter | Receive count of PIN code entry tries
*Class3RemindersAPI* | [**RemindPanCvv**](docs/Class3RemindersAPI.md#remindpancvv) | **Get** /v1/api/cards/{ppan}/pan-cvv | Remind card number
*Class3RemindersAPI* | [**RemindPin**](docs/Class3RemindersAPI.md#remindpin) | **Get** /v1/api/cards/{ppan}/pin | Remind PIN code
*Class3RemindersAPI* | [**ResetPinTryCounter**](docs/Class3RemindersAPI.md#resetpintrycounter) | **Post** /v1/api/cards/{ppan}/pin-try-counter | Reset count of PIN code entry tries
*Class4TransactionsAPI* | [**CardTransactions**](docs/Class4TransactionsAPI.md#cardtransactions) | **Get** /v1/api/cards/{ppan}/transactions | Get card transactions
*Class4TransactionsAPI* | [**DoTransaction**](docs/Class4TransactionsAPI.md#dotransaction) | **Post** /v1/api/cards/{ppan}/transactions | Process card transaction
*Class5AuthorizationsAPI* | [**GetHoldsList**](docs/Class5AuthorizationsAPI.md#getholdslist) | **Get** /v1/api/cards/{ppan}/holds | Get card authorizations
*Class6ClientsAPI* | [**ChangeDeliveryAddress**](docs/Class6ClientsAPI.md#changedeliveryaddress) | **Patch** /v1/api/cards/{ppan}/delivery-address | Change clients delivery address
*Class6ClientsAPI* | [**GetClient**](docs/Class6ClientsAPI.md#getclient) | **Get** /v1/api/clients/{clientId} | Search client by ID
*Class6ClientsAPI* | [**GetClients**](docs/Class6ClientsAPI.md#getclients) | **Get** /v1/api/clients | Get client list
*Class6ClientsAPI* | [**UpdateClientAddress**](docs/Class6ClientsAPI.md#updateclientaddress) | **Patch** /v1/api/clients/{clientId}/address | Change client address fields
*Class6ClientsAPI* | [**UpdateEmail**](docs/Class6ClientsAPI.md#updateemail) | **Put** /v2/api/clients/{clientId}/email | Change client email address
*Class6ClientsAPI* | [**UpdatePassphrase**](docs/Class6ClientsAPI.md#updatepassphrase) | **Put** /v1/api/cards/{ppan}/passphrase | Change card passphrase
*Class6ClientsAPI* | [**UpdatePhone**](docs/Class6ClientsAPI.md#updatephone) | **Put** /v1/api/clients/{clientId}/phone | Change clients phone number
*Class7TokenAPI* | [**ChangeTokenStatus**](docs/Class7TokenAPI.md#changetokenstatus) | **Patch** /v1/api/tokens/{tokenNumber}/state | Change tokens status
*Class7TokenAPI* | [**GetApplePayReferenceData**](docs/Class7TokenAPI.md#getapplepayreferencedata) | **Get** /v1/api/tokens/{ppan}/applepay-reference-data | Get ApplePay reference data
*Class7TokenAPI* | [**GetGooglePayReferenceData**](docs/Class7TokenAPI.md#getgooglepayreferencedata) | **Get** /v1/api/tokens/{ppan}/googlepay-reference-data | Get GooglePay reference data
*Class7TokenAPI* | [**GetTokens**](docs/Class7TokenAPI.md#gettokens) | **Get** /v1/api/tokens/{ppan} | Get card token
*Class8AccountsAPI* | [**GetTransactionsQueueInfo**](docs/Class8AccountsAPI.md#gettransactionsqueueinfo) | **Get** /v1/api/accounts/{cardAccount}/transactions-queue | Get account transactions and commissions in queue
*Class9LimitsAPI* | [**EditLimits**](docs/Class9LimitsAPI.md#editlimits) | **Patch** /v1/api/limits/{ppan} | Change limits data by card
*Class9LimitsAPI* | [**GetLimits**](docs/Class9LimitsAPI.md#getlimits) | **Get** /v1/api/limits/{ppan} | Get limits data by card
*Class9LimitsAPI* | [**GetLimitsRefData**](docs/Class9LimitsAPI.md#getlimitsrefdata) | **Get** /v1/api/limits | Get limits reference data


## Documentation For Models

 - [Address](docs/Address.md)
 - [AddressCreateCard](docs/AddressCreateCard.md)
 - [Amount](docs/Amount.md)
 - [ApiError](docs/ApiError.md)
 - [ApiErrorV2](docs/ApiErrorV2.md)
 - [ApiErrorsV2](docs/ApiErrorsV2.md)
 - [Attachments](docs/Attachments.md)
 - [BatchCardDeliveryAddressApiDto](docs/BatchCardDeliveryAddressApiDto.md)
 - [BatchCardPreferencesApiDto](docs/BatchCardPreferencesApiDto.md)
 - [BatchCardPrivateClientApiDto](docs/BatchCardPrivateClientApiDto.md)
 - [BatchCreateCard](docs/BatchCreateCard.md)
 - [BatchCreateCardTemplateApiDto](docs/BatchCreateCardTemplateApiDto.md)
 - [Beneficiary](docs/Beneficiary.md)
 - [BusinessClient](docs/BusinessClient.md)
 - [BusinessClientCreateCard](docs/BusinessClientCreateCard.md)
 - [BusinessClientInfo](docs/BusinessClientInfo.md)
 - [C2PAddressCreateCard](docs/C2PAddressCreateCard.md)
 - [CardAccountCurrencyInfo](docs/CardAccountCurrencyInfo.md)
 - [CardAccountInfo](docs/CardAccountInfo.md)
 - [CardData1ApiDto](docs/CardData1ApiDto.md)
 - [CardData2ApiDto](docs/CardData2ApiDto.md)
 - [CardData3](docs/CardData3.md)
 - [CardData3HolderApiDto](docs/CardData3HolderApiDto.md)
 - [CardData4](docs/CardData4.md)
 - [CardData6ApiDto](docs/CardData6ApiDto.md)
 - [CardInfo](docs/CardInfo.md)
 - [CardInfoDataArray](docs/CardInfoDataArray.md)
 - [CardOrder](docs/CardOrder.md)
 - [CardPreferences](docs/CardPreferences.md)
 - [CardPreferencesCreateCard](docs/CardPreferencesCreateCard.md)
 - [CardPreferencesGiftCardApiDto](docs/CardPreferencesGiftCardApiDto.md)
 - [CardRenewInfo](docs/CardRenewInfo.md)
 - [CardUserDefinedFields](docs/CardUserDefinedFields.md)
 - [Cardholder](docs/Cardholder.md)
 - [CardholderCreateCard](docs/CardholderCreateCard.md)
 - [ClickToPayCard](docs/ClickToPayCard.md)
 - [ClickToPayCrudApiDto](docs/ClickToPayCrudApiDto.md)
 - [ClickToPayGetData](docs/ClickToPayGetData.md)
 - [ClickToPayProfileApiDto](docs/ClickToPayProfileApiDto.md)
 - [ClickToPayProfileCard](docs/ClickToPayProfileCard.md)
 - [ClientAddress](docs/ClientAddress.md)
 - [ClientInfo](docs/ClientInfo.md)
 - [ClientInfoDataArray](docs/ClientInfoDataArray.md)
 - [Contact](docs/Contact.md)
 - [ContactCreateCard](docs/ContactCreateCard.md)
 - [CreateCard](docs/CreateCard.md)
 - [CreateCardResponse](docs/CreateCardResponse.md)
 - [Data](docs/Data.md)
 - [Data4](docs/Data4.md)
 - [DataArray](docs/DataArray.md)
 - [DateFilter](docs/DateFilter.md)
 - [DateTimeFilter](docs/DateTimeFilter.md)
 - [DeliveryAddress](docs/DeliveryAddress.md)
 - [DeliveryAddressCreate](docs/DeliveryAddressCreate.md)
 - [DeliveryAddressGiftCardApiDto](docs/DeliveryAddressGiftCardApiDto.md)
 - [Document](docs/Document.md)
 - [DocumentCreateCard](docs/DocumentCreateCard.md)
 - [EmailValue](docs/EmailValue.md)
 - [Fee](docs/Fee.md)
 - [File](docs/File.md)
 - [GiftCardOrder](docs/GiftCardOrder.md)
 - [GiftOrderCardResponse](docs/GiftOrderCardResponse.md)
 - [HoldInfo](docs/HoldInfo.md)
 - [HoldsDataArray](docs/HoldsDataArray.md)
 - [Limit](docs/Limit.md)
 - [LimitValues](docs/LimitValues.md)
 - [LimitsData](docs/LimitsData.md)
 - [Merchant](docs/Merchant.md)
 - [Order](docs/Order.md)
 - [OrderCardResponse](docs/OrderCardResponse.md)
 - [PanReferenceData](docs/PanReferenceData.md)
 - [PassphraseValue](docs/PassphraseValue.md)
 - [PhoneValue](docs/PhoneValue.md)
 - [PinAssign](docs/PinAssign.md)
 - [PinTryCounterInfo](docs/PinTryCounterInfo.md)
 - [PinTryReset](docs/PinTryReset.md)
 - [PrivateClient](docs/PrivateClient.md)
 - [PrivateClientCreateCard](docs/PrivateClientCreateCard.md)
 - [PrivateClientGiftCardApiDto](docs/PrivateClientGiftCardApiDto.md)
 - [PrivateClientInfo](docs/PrivateClientInfo.md)
 - [RegistrationAddress](docs/RegistrationAddress.md)
 - [RegistrationAddressCreateApiDto](docs/RegistrationAddressCreateApiDto.md)
 - [RegistrationAddressGiftCardApiDto](docs/RegistrationAddressGiftCardApiDto.md)
 - [StatusChange](docs/StatusChange.md)
 - [TemplateApiDto](docs/TemplateApiDto.md)
 - [Token](docs/Token.md)
 - [TokenInfo](docs/TokenInfo.md)
 - [TokenReferenceData](docs/TokenReferenceData.md)
 - [TokenStatusChange](docs/TokenStatusChange.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionInfo](docs/TransactionInfo.md)
 - [TransactionInfoDataArray](docs/TransactionInfoDataArray.md)
 - [TransactionsQueueInfo](docs/TransactionsQueueInfo.md)
 - [TspSecret](docs/TspSecret.md)
 - [UpdateEmail403Response](docs/UpdateEmail403Response.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@decta.com

